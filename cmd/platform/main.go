package main

import (
	"fmt"
	"os"
)

func main()  {
	// Print angry message to use mattermost command directly
	fmt.Println(`
----------------------------------ERROR----------------------------------------------------
The platform binary has been deprecated, please switch to using the new mattermost binary.
The platform binary will be removed in a future version.
-------------------------------------------------------------------------------------------
`)

	// Execute the real MM binary
	args := os.Args
	args[0] = "mattermost"
	args = append(args, "---platform")

	realMattermost := fileutils.FindF
}
