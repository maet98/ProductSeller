package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Ip   string `json:"IP"`
	Port string `json:"PORT"`
}

var Config Configuration

func GetConfig(params ...string) {
	buf, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(buf), &Config)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", Config.Ip)

}

func (conf Configuration) Url() string {
	log.Println(conf.Ip)
	return conf.Ip + ":" + conf.Port
}
