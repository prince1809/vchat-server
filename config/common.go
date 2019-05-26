package config

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/prince1809/vchat-server/model"
	"io"
	"sync"
)

// commonStore enables code sharing between different backing implementations
type commonStore struct {
	emitter

	configLock             sync.RWMutex
	config                 *model.Config
	configWithoutOverrides *model.Config
	environmentOverrides   map[string]interface{}
}

// Get fetches the current, cached configuration.
func (cs *commonStore) Get() *model.Config {
	cs.configLock.RLock()
	defer cs.configLock.RUnlock()
	return cs.config
}

// load updates the current configuration from the given io.ReadCloser.
//
// This function assumes no lock has been acquired, as it acquires a write lock itself.
func (cs *commonStore) load(f io.ReadCloser, needsSave bool, validate func(*model.Config) error, persist func(*model.Config) error) error {
	// Duplicate f so that we can read a configuration without applying environment overrides
	f2 := new(bytes.Buffer)
	tee := io.TeeReader(f, f2)

	allowEnvironmentOverrides := true
	loadedCfg, environmentOverrides, err := unmarshalConfig(tee, allowEnvironmentOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config with env overrides")
	}

	// Keep track of the original values that the Environment settings overrode
	loadedCfgWithoutEnvOverrides, _, err := unmarshalConfig(f2, false)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config without env overrides")
	}

	// setDefaults generates various keys and salts if not previously configured. Determine if
	// such a change will be made before invoking.
	needsSave = needsSave || loadedCfg.SqlSettings.AtRestEncrypt == nil || len(*loadedCfg.SqlSettings.AtRestEncrypt) == 0
	needsSave = needsSave || loadedCfg.FileSettings.PublicLinkSalt == nil || len(*loadedCfg.FileSettings.PublicLinkSalt) == 0

	loadedCfg.SetDefaults()

	if validate != nil {
		if err = validate(loadedCfg); err != nil {
			return errors.Wrap(err, "invalid config")
		}
	}

	if changed := fixConfig(loadedCfg); changed {
		needsSave = true
	}

	cs.configLock.Lock()
	var unlockOnce sync.Once
	defer unlockOnce.Do(cs.configLock.Unlock)

	if needsSave && persist != nil {
		cfgWithoutEnvOverrides := removeEnvOverrides(loadedCfg, loadedCfgWithoutEnvOverrides, environmentOverrides)
		if err = persist(cfgWithoutEnvOverrides); err != nil {
			return errors.Wrap(err, "failed to persist required changes after load")
		}
	}

	oldCfg := cs.config
	cs.config = loadedCfg
	cs.configWithoutOverrides = loadedCfgWithoutEnvOverrides
	cs.environmentOverrides = environmentOverrides

	unlockOnce.Do(cs.configLock.Unlock)

	// Notify listeners synchronously, Ideally this should be asynchronous,but existing code
	// assumes this and there would be increased complexity to avoid racing updates.
	cs.invokeConfigListeners(oldCfg, loadedCfg)

	return nil
}

// validate checks if the given configuration is valid
func (cs *commonStore) validate(cfg *model.Config) error {
	if err := cfg.IsValid(); err != nil {
		return err
	}
	return nil
}
