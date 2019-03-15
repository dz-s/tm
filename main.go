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
	mode := os.FileMode(0700)

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

	os.Mkdir(client, os.FileMode(mode))
	os.OpenFile(filepath.Join(client, string(Client)), os.O_RDONLY|os.O_CREATE, mode)
	os.Mkdir(project, os.FileMode(mode))
	os.OpenFile(filepath.Join(project, string(Project)), os.O_RDONLY|os.O_CREATE, mode)

	fmt.Println("initialize has been done successfully")
	return nil
}

// func add(entity Entity) {
// }

// func remove(entity Entity) {
// }

func main() {

	app := cli.NewApp()

	app.Name = "tm"
	app.Usage = "Minimalistic task management"

	// we create our commands
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initial setup",
			Action: _init,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
