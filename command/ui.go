package command

import (
	"os"

	"github.com/mitchellh/cli"
)

func NewUI() cli.Ui {
	return &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
}
