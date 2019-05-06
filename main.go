package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const commitMsgFileName = "COMMIT_EDITMSG"

func showUsageAndExit() {
	fmt.Fprintln(os.Stderr, "Usage: git-zoo <COMMIT_EDITMSG_FILE>")
	os.Exit(1)
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

func getEditedCommitMessage(filePath string) (string, error) {
	commitMsg := animals() + " "

	fp, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		commitMsg = commitMsg + scanner.Text() + "\n"
	}
	fp.Close()

	return commitMsg, nil
}

func writeCommitMessage(filePath string, msg string) error {
	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()
	writer := bufio.NewWriter(fp)

	_, err = writer.WriteString(msg)
	if err != nil {
		return err
	}
	writer.Flush()
	fp.Close()

	return nil
}

func main() {
	if len(os.Args) <= 1 {
		showUsageAndExit()
	}

	filePath := os.Args[1]
	fileName := filepath.Base(filePath)
	if fileName != commitMsgFileName {
		showUsageAndExit()
	}

	commitMsg, err := getEditedCommitMessage(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = writeCommitMessage(filePath, commitMsg)
	if err != nil {
		log.Fatal(err)
	}
}
