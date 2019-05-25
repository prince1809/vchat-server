package commands

import (
	"github.com/mattermost/viper"
	"github.com/prince1809/vchat-server/app"
	"github.com/prince1809/vchat-server/utils"
	"github.com/spf13/cobra"
)

func InitDBCommandContextCobra(command *cobra.Command) (*app.App, error)  {
	config := viper.GetString("config")

	a, err := InitDBCommandContext(config)

	if err != nil {
		// Returning an error just prints the usage message, so actually panic
		panic(err)
	}

	return a, nil
}

func InitDBCommandContext(configDSN string) (*app.App, error) {
	if err := utils.TranslationsPreInit(); err != nil {
		return nil, err
	}
	//model.App

	//return a, nil
	panic("implement me")
}
