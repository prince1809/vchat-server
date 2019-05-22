package main

import (
	"github.com/prince1809/vchat-server/cmd/mattermost/commands"
	"os"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
