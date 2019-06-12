package main

import (
	"fmt"
	"bytes"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

func excecuteTemplate(title string) string {
	lastLog := getLogStruct()

	lastLog.CreatedDate = time.Now().String()
	lastLog.Title = title
	var tpl bytes.Buffer
	tmpl, err := template.ParseFiles(rootDir + "/Task.md")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tpl, lastLog)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

func isClient() bool {
	_, err := os.Stat(string(Client))
	return (err == nil)
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

		f, err := os.OpenFile(filepath.Join(root, name + ".md"), os.O_RDWR|os.O_CREATE, MODE)
		out := excecuteTemplate(name)
		fmt.Println(out)
		_, err = f.WriteString(out)
		f.Sync()
		if err != nil {
			return err
		}
	}
	os.OpenFile(filepath.Join(root, string(_type)), os.O_RDONLY|os.O_CREATE, MODE) // add entity identifier
	return nil
}
