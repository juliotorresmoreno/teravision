package main

import (
	"fmt"

	"github.com/juliotorresmoreno/teravision/bootstrap"
	"github.com/juliotorresmoreno/teravision/config"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Init()
	conf := config.GetConf()
	fmt.Println(conf.DatabaseDSN)
}
