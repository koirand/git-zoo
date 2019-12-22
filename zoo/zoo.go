package zoo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func EditCommitMsg(path string) error {
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

	fmt.Fprint(f, getAnimal(), " ", string(commitMsg))
	return nil
}

func MakeSymboliclink() error {
	const symlinkPath = ".git/hooks/prepare-commit-msg"
	gitRootDir, err := getGitRootDir()
	if err != nil {
		return err
	}
	binPath, _ := os.Executable()

	cmd := exec.Command("ln", "-s", binPath, filepath.Join(gitRootDir, symlinkPath))
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.New("Error: Failed to link to github hooks.")
	}
	return nil
}

func getAnimal() string {
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

func getGitRootDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stderr = os.Stderr

	gitDirPathByte, err := cmd.Output()
	if err != nil {
		return "", errors.New("Error: Run command in the git repository.")
	}

	return strings.Trim(string(gitDirPathByte), "\n"), nil
}
