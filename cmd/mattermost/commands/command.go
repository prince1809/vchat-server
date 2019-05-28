package commands

import "github.com/spf13/cobra"

var CommandCmd = &cobra.Command{
	Use:   "command",
	Short: "Management of slash commands",
}

var CommandCreateCmd = &cobra.Command{
	Use:     "Create [team]",
	Short:   "Create a custom slash command",
	Long:    "Create a custom slash command for the specified team.",
	Args:    cobra.MinimumNArgs(1),
	Example: `  command create myteam --title MyCommand --description "My Command Description" --trigger-word mycommand --url http://localhost:8000/my-slash-handler --creator myusername --response-username my-bot-username --icon http://localhost:8000/my-slash-handler-bot-icon.png --autocomplete --post`,
	RunE:    createCommandCmdF,
}

func createCommandCmdF(command *cobra.Command, args []string) error {
	a, err := InitDBCommandContextCobra(command)
	if err != nil {
		return err
	}
	defer a.Shutdown()
	return nil
}
