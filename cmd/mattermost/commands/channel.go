package commands

import "github.com/spf13/cobra"

var ChannelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Management of channels",
}

var ChannelCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a channel",
	Long:  `Create a channel`,
	Example: `  channel create --team myteam --name mynewchannel --display_name "My New Channel"
	channel create --team myteam --name mynewprivatechannel --display_name "My New Private Channel" --private`,
	RunE: createCommandCmdF,
}

func init() {
	ChannelCreateCmd.Flags().String("name", "", "Channel Name")

	RootCmd.AddCommand(ChannelCmd)
}
