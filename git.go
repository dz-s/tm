package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

func getLastCommit() string {
	cmd := exec.Command("git", "rev-parse", "--verify", "HEAD") // git rev-parse --verify HEAD
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func getAuthor(c *cli.Context) error {
	commitID := getLastCommit()
	cmd := exec.Command("git", "show", "--format", "%aN <%aE> ", commitID)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
	return nil
}
