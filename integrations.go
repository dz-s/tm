package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml" //TODO: change to searching for particular key instead of parsing
	"github.com/dz-s/tm/trello"
)

func getCredentials(key string, token string) (string, string) { //Might be turned into SyncTrello

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

	return m[key], m[token]
}

func getMetadata(keyName string, tokenName string, baseURL string) {

	key, token := getCredentials(keyName, tokenName)
	res, err := http.Get(
		"https://api.trello.com/1/boards/5bab58ffb5f96804c74f737e/cards?fields=all&key=" +
			string(key) +
			"&token=" +
			string(token))

	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var tasks trello.Tasks
	err = json.Unmarshal(buf, &tasks)
	fmt.Println(tasks[0].Desc)

}
