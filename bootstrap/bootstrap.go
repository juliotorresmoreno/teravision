package bootstrap

import "github.com/juliotorresmoreno/teravision/config"

// Init init the app
func Init() {
	config.Init("env.yml")
}
