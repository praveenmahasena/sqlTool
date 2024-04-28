package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

func readConfig() (*YamlConfig, error) {
	var path, pathErr = os.Getwd()
	if pathErr != nil {
		return nil, pathErr
	}

	var file, fileErr = os.ReadFile(path + "/psql.yaml")
	if fileErr != nil {
		return nil, fileErr
	}

	var config = new(YamlConfig)
	var err = yaml.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
