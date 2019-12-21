package command

import (
	"fmt"
	"path/filepath"

	"github.com/koirand/git-zoo/tools"
	"github.com/mitchellh/cli"
)

type UninstallCommand struct {
	Ui cli.Ui
}

func NewUninstallCommand() (cli.Command, error) {
	return new(UninstallCommand), nil
}

func (e *UninstallCommand) Help() string {
	return "git-zoo uninstall"
}

func (e *UninstallCommand) Synopsis() string {
	return "Uninstall git-zoo from current git repository"
}

func (e *UninstallCommand) Run(args []string) int {
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
