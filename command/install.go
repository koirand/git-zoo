package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/koirand/git-zoo/tools"
	"github.com/mitchellh/cli"
)

type InstallCommand struct {
	Ui cli.Ui
}

func NewInstallCommand() (cli.Command, error) {
	return new(InstallCommand), nil
}

func (e *InstallCommand) Help() string {
	return "git-zoo install"
}

func (e *InstallCommand) Synopsis() string {
	return "Install git-zoo to current git repository"
}

func (e *InstallCommand) Run(args []string) int {
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
