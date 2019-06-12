package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

type GitLogItem struct {
	Commit     string `json:"commit"`
	Author     string `json:"author"`
	Date	string `json:"date"`
	CreatedDate	string 
	Title	string 
}

func getLastCommit() string {
	cmd := exec.Command("git", "show", "--name-status") // git rev-parse --verify HEAD
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func getAuthor(c *cli.Context) error {
	getLogStruct()
	return nil
}

func getLogStruct() GitLogItem {
	cmd := exec.Command("/bin/sh", rootDir + "/parseGitLog.sh")
	res, err := cmd.Output()
	var lastLog GitLogItem
	err = json.Unmarshal([]byte(res), &lastLog)
	if err != nil {
		panic(err)
	}
	return lastLog
}
