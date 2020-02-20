package models

import (
	"log"

	"github.com/juliotorresmoreno/teravision/db"
)

// Migrate update all shema in bd
func Migrate() {
	conn, err := db.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Sync2(User{})
	if err != nil {
		log.Fatal(err)
	}
}
