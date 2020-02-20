package db_test

import (
	"testing"

	"github.com/juliotorresmoreno/teravision/config"
	"github.com/juliotorresmoreno/teravision/db"
)

func TestNewConn(t *testing.T) {
	config.Init("../env.yml")
	conn, err := db.NewConn()
	if err != nil {
		t.Error(err)
		return
	}
	n, err := conn.SQL("select 1").Count()
	if err != nil {
		t.Error("Connection failure")
		return
	}
	if n != 1 {
		t.Error("Connection failure")
		return
	}
	defer conn.Close()
}
