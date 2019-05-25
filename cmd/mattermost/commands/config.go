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

func configValidateCmdF(command *cobra.Command, args []string) error {
	utils.TranslationsPreInit()
	model.AppErrorInit(utils.T)

	_, err := getConfigStore(command)
}

func getConfigStore(command *cobra.Command) (config.Store, error) {
	if err := utils.TranslationsPreInit(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize i18n")
	}

	configDSN := viper.GetString("config")

	configStore, err := config.Ne
}
