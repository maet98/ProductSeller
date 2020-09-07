package main

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ip   string
	port string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	filename := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(filename, &configuration)
	return configuration
}
