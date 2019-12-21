package command

import (
	"fmt"
	"path/filepath"

	"github.com/koirand/git-zoo/tools"
	"github.com/mitchellh/cli"
)

type DisableCommand struct {
	Ui cli.Ui
}

func NewDisableCommand() (cli.Command, error) {
	return new(DisableCommand), nil
}

func (e *DisableCommand) Help() string {
	return "git-zoo disable"
}

func (e *DisableCommand) Synopsis() string {
	return "Disable git-zoo in current git repository"
}

func (e *DisableCommand) Run(args []string) int {
	gitRootDir, err := tools.GetGitRootDir()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if err := tools.Unlink(
		filepath.Join(string(gitRootDir), ".git/hooks/prepare-commit-msg"),
	); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
