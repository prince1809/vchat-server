package config

import (
	"bytes"
	"encoding/json"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/viper"
	"github.com/pkg/errors"
	"github.com/prince1809/vchat-server/model"
	"github.com/prince1809/vchat-server/utils/jsonutils"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
)

// newViper creates an instance of viper.Viper configured fro parsing a configuration.
func newViper(allowEnvironmentOverrides bool) *viper.Viper {
	v := viper.New()

	v.SetConfigType("json")

	v.AllowEmptyEnv(true)

	if allowEnvironmentOverrides {
		v.SetEnvPrefix("mm")
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		v.AutomaticEnv()
	}

	// set zeroed defaults fro all the config settings so that Viper knows what environment variables
	// it needs to be looking for. The correct defaults will later be applied using Config.SetDefaults.
	defaults := getDefaultsFromStruct(model.Config{})

	for key, value := range defaults {
		if key == "PluginSettings.Plugins" || key == "PluginSettings.PluginStates" {
			continue
		}

		v.SetDefault(key, value)
	}

	return v
}

func getDefaultsFromStruct(s interface{}) map[string]interface{} {
	return flattenStructToMap(structToMap(reflect.TypeOf(s)))
}

// converts a struct type into a nested map with keys mathing the struct's fields and values
// matching the zeroed value of the corresponding field.
func structToMap(t reflect.Type) (out map[string]interface{}) {
	defer func() {
		if r := recover(); r != nil {
			mlog.Error("Panicked in structToMap. This could never happen.", mlog.Any("err", r))
		}
	}()

	if t.Kind() != reflect.Struct {
		// Should never hit this, but this will prevent a panic if that does happen somewhow
		return nil
	}

	out = map[string]interface{}{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		var value interface{}

		switch field.Type.Kind() {
		case reflect.Struct:
			value = structToMap(field.Type)
		case reflect.Ptr:
			indirectType := field.Type.Elem()

			if indirectType.Kind() == reflect.Struct {
				// Follow pointers to structs since we need to define defaults for their fields
				value = structToMap(indirectType)
			} else {
				value = nil
			}
		default:
			value = reflect.Zero(field.Type).Interface()
		}
		out[field.Name] = value
	}

	return
}

// Flattens a nested map so that the result is a single map with keys corresponding to the
// path through the original map. For example,
// {
// 		"a": {
// 			"b": 1
// 		},
//		"c": "sea"
// }
// would flatten to
// {
// {
//		"a.b": 1
// 		"c": "sea"
//	}

func flattenStructToMap(in map[string]interface{}) map[string]interface{} {
	return in
}

// marshalConfig converts the given configuration into JSON bytes for persistence.
func marshalConfig(cfg *model.Config) ([]byte, error) {
	return json.MarshalIndent(cfg, "", "  ")
}

// unmarshalConfig unmarshals a raw configuration into a Config model and environment variable overrides.
func unmarshalConfig(r io.Reader, allowEnvironmentOverrides bool) (*model.Config, map[string]interface{}, error) {
	// Pre-flight check the syntax of the configuration file to improve error messaging.
	configData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to read")
	}

	var rawConfig interface{}
	if err = json.Unmarshal(configData, &rawConfig); err != nil {
		return nil, nil, jsonutils.HumanizeJsonError(err, configData)
	}

	v := newViper(allowEnvironmentOverrides)
	if err := v.ReadConfig(bytes.NewReader(configData)); err != nil {
		return nil, nil, err
	}

	var config model.Config
	unmarshalErr := v.Unmarshal(&config)
	if unmarshalErr == nil {
		config.PluginSettings.Plugins = make(map[string]map[string]interface{})
	}

	envConfig := v.EnvSettings()

	return &config, envConfig, unmarshalErr
}
