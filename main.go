package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Entity string

var tok string
var proj string

const (
	Client  Entity = ".client"
	Project Entity = ".project"
	Task    Entity = ".task"
)

/*_init ...
Initialize client and project directives
*/
func _init(c *cli.Context) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.Wrap(err, "initialize failed (dirname error)")
	}

	client := filepath.Join(dir, c.Args().Get(0))

	project := filepath.Join(client, c.Args().Get(1))
	/**
	+-----+---+--------------------------+
	| rwx | 7 | Read, write and execute  |
	| rw- | 6 | Read, write              |
	| r-x | 5 | Read, and execute        |
	| r-- | 4 | Read,                    |
	| -wx | 3 | Write and execute        |
	| -w- | 2 | Write                    |
	| --x | 1 | Execute                  |
	| --- | 0 | no permissions           |
	+------------------------------------+

	+------------+------+-------+
	| Permission | Octal| Field |
	+------------+------+-------+
	| rwx------  | 0700 | User  |
	| ---rwx---  | 0070 | Group |
	| ------rwx  | 0007 | Other |
	+------------+------+-------+
	*/

	os.Mkdir(client, MODE)
	os.OpenFile(filepath.Join(client, string(Client)), os.O_RDONLY|os.O_CREATE, MODE)
	os.Mkdir(project, MODE)
	os.OpenFile(filepath.Join(project, string(Project)), os.O_RDONLY|os.O_CREATE, MODE)

	fmt.Println("initialize has been done successfully")
	return nil
}

func isClient() bool {
	_, err := os.Stat(string(Client))
	if err != nil {
		return false
	}
	return true
}

func isProject() bool {
	_, err := os.Stat(string(Project))
	if err != nil {
		return false
	}
	return true
}

func isTask() bool {
	_, err := os.Stat(string(Task))
	if err != nil {
		return false
	}
	return true
}

func createEntity(name string, _type Entity) error {
	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "Couldn't get root dir path")
	}
	os.Mkdir(filepath.Join(dir, name), MODE)
	root := dir + "/" + name
	if _type == Task {
		os.OpenFile(filepath.Join(root, name), os.O_RDWR|os.O_CREATE, MODE)
	}
	os.OpenFile(filepath.Join(root, string(_type)), os.O_RDONLY|os.O_CREATE, MODE) // add entity identifier
	return nil
}

func add(c *cli.Context) {
	name := c.Args().Get(0)
	if isClient() {
		createEntity(name, Project)
	} else if isProject() {
		createEntity(name, Task)
	} else {
		createEntity(name, Task)
	}
}

func _import(c *cli.Context) {
	fmt.Println(&tok)
	fmt.Println(importProject(tok, proj))
}

func main() {

	app := cli.NewApp()

	app.Name = "tm"
	app.Usage = "Minimalistic task management"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "tok",
			Value:       "key",
			Usage:       "access token for board API",
			Destination: &tok,
		},
		cli.StringFlag{
			Name:        "proj",
			Value:       "foobar",
			Usage:       "project name",
			Destination: &proj,
		},
	}
	// we create our commands
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initial setup",
			Action: _init,
		},
		{
			Name:   "add",
			Usage:  "Initial setup",
			Action: add,
		},
		{
			Name:   "last",
			Usage:  "Initial setup",
			Action: getAuthor,
		},
		{
			Name:   "import",
			Usage:  "Import your project from board",
			Action: _import,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
