package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"github.com/gosimple/slug"
	"log"
)

var key, tok, proj string
var title string
var rootDir string

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

func add(c *cli.Context) {

	name := slug.Make(title)

	if isClient() {
		createEntity(name, Project)
	} else if isProject() {
		createEntity(name, Task)
	} else {
		createEntity(name, Task)
	}
}

func _import(c *cli.Context) {
	// fmt.Println(key, tok, proj)
	// key, token := getMetadata("trello_Key", "trello_Token", "testProject")
	// fmt.Println(getCredentials(key, tok))
}

func main() {

	rootDir = "~/tm"
	fmt.Println(rootDir)
	//_, templates := template.ParseFiles("Task.md")
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
		{
			Name:   "add",
			Usage:  "Add entity",
			Action: add,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "t",
					Value:       "title",
					Usage:       "title for Task.md template",
					Destination: &title,
				},
			},
		},
		{
			Name:   "last",
			Usage:  "Get last change",
			Action: getAuthor,
		},
		{
			Name:   "import",
			Usage:  "Import your project from board",
			Action: _import,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "key",
					Value:       "key",
					Usage:       "access key for board API",
					Destination: &key,
				},
				cli.StringFlag{
					Name:        "tok",
					Value:       "tok",
					Usage:       "access token for board API",
					Destination: &tok,
				},
				cli.StringFlag{
					Name:        "proj",
					Value:       "foobar",
					Usage:       "project name",
					Destination: &proj,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
