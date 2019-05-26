package config

import "github.com/prince1809/vchat-server/model"

// removeEnvOverrides returns a new config without the given environment overrides.
// If a config variable has an environment override, that variable is set to the value that was
// read from the store.
func removeEnvOverrides(cfg, cfgWithoutEnv *model.Config, envOverrides map[string]interface{}) *model.Config {

	return cfg
}
