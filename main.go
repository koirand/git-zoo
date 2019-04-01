package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

const commitMsgFileName = "COMMIT_EDITMSG"

func main() {
	app := cli.NewApp()

	app.Name = "git-zoo"
	app.Usage = "add animals emoji to git commit message."
	app.Version = "0.1.0"

	app.Action = func(c *cli.Context) error {
		filePath := c.Args().Get(0)
		fileName := filepath.Base(filePath)

		if c.NArg() == 0 || fileName != commitMsgFileName {
			return cli.NewExitError("Requires an argument that is path to COMMIT_EDITMSG", 1)
		}

		// read commit message
		fp, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(fp)

		// add emoji to commit message
		commitMsg := animals() + " "
		for scanner.Scan() {
			commitMsg = commitMsg + scanner.Text() + "\n"
		}
		fmt.Println(commitMsg)
		fp.Close()

		// write commit message
		fp, err = os.Create(filePath)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer fp.Close()
		writer := bufio.NewWriter(fp)

		_, err = writer.WriteString(commitMsg)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		writer.Flush()
		fp.Close()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
