package model

type PluginOption struct {
	// The display name for the option.
	DisplayName string `json:"display_name" yaml:"display_name"`

	// The string value for the option.
	Value string `json:"value" yaml:"value"`
}

type PluginSetting struct {
	// The key that the setting will be assigned to in the configuration file.
	Key string `json:"key"`

	// The display name for the setting.
	DisplayName string `json:"display_name" yaml:"display_name"`

	// The type of the setting.
	//
	// "bool" will result in a boolean true or false setting.
	//
	// "dropdown" will result in a string setting that allows the user to select from a list of
	// pre-defined options.
	//
	// "generated" will result in a string setting that is set to a random, cryptographically secure
	// string.
	//
	// "radio" will result in a string setting that allows the user to select from a short selection
	// of pre-defined options.
	//
	// "text" will result in a string setting that can be typed in manually.
	//
	// "longtext" will result in a multi line string that can be typed in manually.
	//
	// "username" will result in a text setting that will autocomplete to a username.
	Type string `json:"type" yaml:"type"`

	// The help text to display to the user.
	HelpText string `json:"help_text" yaml:"help_text"`


}

type Manifest struct {
}
