package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"fmt"
	"os"
)

type Conf struct {
	Youtube struct {
			ApiKey string `yaml:"apiKey"`
		}
}

var Configs Conf

func initConfig() {
	absPath, err := filepath.Abs("./config.yml")
	CheckErr(err)
	data, err := ioutil.ReadFile(absPath)
	CheckErr(err)
	err = yaml.Unmarshal(data, &Configs)
	CheckErr(err)
	if Configs.Youtube.ApiKey == "{PAST YOUR KEY HERE}" {
		fmt.Println("ERROR: Wrong Youtube API key!!!")
		os.Exit(1)
	}
}