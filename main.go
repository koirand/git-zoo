package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/koirand/git-zoo/zoo"
)

func main() {
	if len(os.Args) > 1 && filepath.Base(os.Args[1]) == "COMMIT_EDITMSG" {
		// When -c, -C, or -amend are used as git commit options, do nothing
		if len(os.Args) > 2 && os.Args[2] == "commit" {
			os.Exit(0)
		}

		// Edit commit message
		if err := zoo.EditCommitMsg(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Install git hooks
		if err := zoo.MakeSymboliclink(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
