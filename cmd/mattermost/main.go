package main

import (
	"fmt"
	"github.com/prince1809/vchat-server/cmd/mattermost/commands"
	"os"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		fmt.Println("Error starting:", err)
		os.Exit(1)
	}
}
