package main

import (
	"os"
	"path/filepath"
	
	"github.com/pkg/errors"
)

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
		os.OpenFile(filepath.Join(root, name), os.O_RDWR|os.O_CREATE, MODE)
	}
	os.OpenFile(filepath.Join(root, string(_type)), os.O_RDONLY|os.O_CREATE, MODE) // add entity identifier
	return nil
}
