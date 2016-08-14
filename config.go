package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
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
}