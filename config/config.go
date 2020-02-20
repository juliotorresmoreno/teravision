package config

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

// Config all data of configuration
type Config struct {
	DriverName  string `yaml:"driver_name"`
	DatabaseDSN string `yaml:"database_dsn"`
}

var c Config

/*
func init() {
	Init("env.yml")
}*/

// Init default config
func Init(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(f, &c)
}

// GetConf Returns config
func GetConf() Config {
	return c
}
