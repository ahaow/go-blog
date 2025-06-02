package core

import (
	"go-blog/config"
	"go-blog/utils"
	"log"

	"gopkg.in/yaml.v3"
)

func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYaml()
	if err != nil {
		log.Fatalf("failed to load configuration:%v", err)
	}
	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("failed to load configuration:%v", err)
	}
	return c
}
