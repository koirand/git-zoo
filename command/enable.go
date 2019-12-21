package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/koirand/git-zoo/tools"
	"github.com/mitchellh/cli"
)

type EnableCommand struct {
	Ui cli.Ui
}

func NewEnableCommand() (cli.Command, error) {
	return new(EnableCommand), nil
}

func (e *EnableCommand) Help() string {
	return "git-zoo enable"
}

func (e *EnableCommand) Synopsis() string {
	return "Enable git-zoo in current git repository"
}

func (e *EnableCommand) Run(args []string) int {
	gitRootDir, err := tools.GetGitRootDir()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	binPath, _ := os.Executable()
	if err := tools.MakeSymboliclink(
		binPath,
		filepath.Join(string(gitRootDir), ".git/hooks/prepare-commit-msg"),
	); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
