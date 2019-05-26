package commands

import (
	"github.com/mattermost/viper"
	"github.com/pkg/errors"
	"github.com/prince1809/vchat-server/config"
	"github.com/prince1809/vchat-server/model"
	"github.com/prince1809/vchat-server/utils"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration",
}

var ValidateConfigCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate config file",
	Long:  "If the config file is valid, this command will output a success message and have a zero exit code. If it is invalid, this command will output an error and have non-zero exit code.",
	RunE:  configValidateCmdF,
}

var ConfigSubpathCmd = &cobra.Command{
	Use:   "subpath",
	Short: "Update client asset loading to use the configured subpath",
	Long:  "Update the hard-coded production client asset paths to take into account Mattermost running on a subpath.",
	Example: `  config subpath
	config subpath --path /mattermost
	config subpath --path /`,
	RunE: configSubpathCmdF,
}

func init() {
	ConfigCmd.AddCommand(
		ValidateConfigCmd,
		ConfigSubpathCmd,
	)
	RootCmd.AddCommand(ConfigCmd)
}

func configValidateCmdF(command *cobra.Command, args []string) error {
	utils.TranslationsPreInit()
	model.AppErrorInit(utils.T)

	_, err := getConfigStore(command)
	if err != nil {
		return err
	}

	CommandPrettyPrintln("The document is valid")
	return nil
}

func configSubpathCmdF(commmand *cobra.Command, args []string) error {
	return nil
}

func getConfigStore(command *cobra.Command) (config.Store, error) {
	if err := utils.TranslationsPreInit(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize i18n")
	}

	configDSN := viper.GetString("config")

	configStore, err := config.NewStore(configDSN, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize config store")
	}

	return configStore, nil
}
