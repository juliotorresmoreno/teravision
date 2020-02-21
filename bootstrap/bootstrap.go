package bootstrap

import (
	"github.com/juliotorresmoreno/teravision/config"
	"github.com/juliotorresmoreno/teravision/helpers"
	"github.com/juliotorresmoreno/teravision/models"
)

// Init init the app
func Init() {
	config.Init("env.yml")
	helpers.Init()
	models.Migrate()
}
