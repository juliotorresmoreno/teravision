package models

import "time"

// User model of users
type User struct {
	ID   uint64    `xorm:"id integer pk"`
	Name string    `xorm:"name varchar(100) not null"`  // ○ Nombre Completo (Requerido)
	DNI  string    `xorm:"dni varchar(20) not null uk"` // ○ DNI o cédula (Requerido)
	date time.Time `xorm:"date time with time zone"`    // ○ Fecha de Nacimiento. (El usuario puede o no enviar esta información)
}

// TableName name in bd
func (User) TableName() string {
	return "users"
}
