package main

import (
	"github.com/juliotorresmoreno/teravision/bootstrap"
	"github.com/juliotorresmoreno/teravision/server"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Init()
	svr := server.NewServer()
	svr.Listen()
}
