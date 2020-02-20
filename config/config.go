package config

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

// Config all data of configuration
type Config struct {
	DatabaseDSN string `yaml:"database_dsn"`
}

var c Config

func init() {
	f, err := ioutil.ReadFile("env.yml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(f, &c)
}

// GetConf Returns config
func GetConf() Config {
	return c
}
