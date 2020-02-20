package main

import (
	"fmt"

	"github.com/juliotorresmoreno/teravision/config"
)

func main() {
	conf := config.GetConf()
	fmt.Println(conf.DatabaseDSN)
}
