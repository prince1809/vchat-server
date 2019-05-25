package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/utils/fileutils"
	"os"
	"syscall"
)

func main() {
	// Print angry message to use mattermost command directly
	fmt.Println(`
----------------------------------ERROR----------------------------------------------------
The platform binary has been deprecated, please switch to using the new mattermost binary.
The platform binary will be removed in a future version.
-------------------------------------------------------------------------------------------`)

	// Execute the real MM binary
	args := os.Args
	args[0] = "mattermost"
	args = append(args, "---platform")

	realMattermost := fileutils.FindFile("mattermost")
	if realMattermost == "" {
		realMattermost = fileutils.FindFile("bin/mattermost")
	}

	if realMattermost == "" {
		fmt.Println("Could not start Mattermost, use the mattermost command directly:")
	} else if err := syscall.Exec(realMattermost, args, nil); err != nil {
		fmt.Printf("Could not start Mattermost, use the mattermost command directly: %s\n", err.Error())
	}
}
