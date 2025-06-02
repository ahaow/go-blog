package utils

import (
	"go-blog/global"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

func LoadYaml() ([]byte, error) {
	return os.ReadFile(configFile)
}

func SaveYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, byteData, fs.ModePerm)
}
