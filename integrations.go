package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml" //TODO: change to searching for particular key instead of parsing
)

func importProject(key string, projectName string) string { //Might be turned into SyncTrello

	if _, err := os.Stat(".tmrc"); os.IsNotExist(err) {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadFile(".tmrc")
	if err != nil {
		log.Fatal(err)
	}

	str := string(buf)

	m := make(map[string]string)

	if _, err := toml.Decode(str, &m); err != nil {
		log.Fatal(err)
	}

	return m[key]
}
