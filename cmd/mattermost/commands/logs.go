package commands

import (
	"github.com/prince1809/vchat-server/mlog/human"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Display logs in human-readable format",
	RunE:  logsCmdF,
}

func init() {
	LogsCmd.Flags().Bool("logrus", false, "Use logrus for formatting.")
	RootCmd.AddCommand(LogsCmd)
}

func logsCmdF(command *cobra.Command, args []string) error {
	// check stdin to see if we have a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	var input io.Reader
	if fi.Size() == 0 && fi.Mode()&os.ModeNamedPipe == 0 {
		file, err := os.Open("mattermost.log")
		if err != nil {
			return err
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}
	var writer human.LogWriter

	if flag, _ := command.Flags().GetBool("logrus"); flag {
		writer = human.NewLogrusWriter(os.Stdout)
	} else {
		writer = human.NewLogrusWriter(os.Stdout)
	}
	human.ProcessLogs(input, writer)

	return nil
}
