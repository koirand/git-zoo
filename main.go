package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/koirand/git-zoo/command"
	"github.com/mitchellh/cli"
)

func main() {
	if len(os.Args) > 1 && filepath.Base(os.Args[1]) == "COMMIT_EDITMSG" {
		// When -c, -C, or -amend are used as git commit options,
		// Do not change the commit message.
		if len(os.Args) > 2 && os.Args[2] == "commit" {
			os.Exit(0)
		}

		if err := edit(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	c := cli.NewCLI("git-zoo", "0.1.0")
	c.Args = os.Args[1:]
	ui := command.NewUI()
	c.Commands = map[string]cli.CommandFactory{
		"install": func() (cli.Command, error) {
			return &command.InstallCommand{Ui: ui}, nil
		},
		"uninstall": func() (cli.Command, error) {
			return &command.UninstallCommand{Ui: ui}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(err.Error())
	}

	os.Exit(exitStatus)
}

func edit(path string) error {
	commitMsg, err := (func(path string) ([]byte, error) {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return ioutil.ReadAll(f)
	})(path)

	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprint(f, animals(), " ", string(commitMsg))
	return nil
}

func animals() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(14)
	switch index {
	case 0:
		return "🐶"
	case 1:
		return "🐱"
	case 2:
		return "🐭"
	case 3:
		return "🐹"
	case 4:
		return "🐰"
	case 5:
		return "🦊"
	case 6:
		return "🐻"
	case 7:
		return "🐼"
	case 8:
		return "🐨"
	case 9:
		return "🐯"
	case 10:
		return "🦁"
	case 11:
		return "🐮"
	case 12:
		return "🐷"
	case 13:
		return "🐸"
	case 14:
		return "🐵"
	}
	return ""
}
