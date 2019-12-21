package tools

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func GetGitRootDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stderr = os.Stderr

	gitDirPathByte, err := cmd.Output()
	if err != nil {
		return "", errors.New("Error: Run command in the git repository.")
	}

	return strings.Trim(string(gitDirPathByte), "\n"), nil
}

func MakeSymboliclink(source string, target string) error {
	cmd := exec.Command("ln", "-s", source, target)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.New("Error: Failed to link to github hooks.")
	}
	return nil
}

func Unlink(path string) error {
	cmd := exec.Command("unlink", path)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.New("Error: Failed to unlink github hooks.")
	}
	return nil
}
