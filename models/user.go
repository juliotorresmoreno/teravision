package models

import "time"

// User model of users
type User struct {
	ID   uint      `json:"id" xorm:"bigint not null autoincr pk"`
	Name string    `json:"name" xorm:"name varchar(100) not null"`      // ○ Nombre Completo (Requerido)
	DNI  string    `json:"dni"  xorm:"dni varchar(20) not null unique"` // ○ DNI o cédula (Requerido)
	Date time.Time `json:"date" xorm:"date time with time zone"`        // ○ Fecha de Nacimiento. (El usuario puede o no enviar esta información)
}

// CleanForInsert clean for insert
func (that *User) CleanForInsert() {
	that.ID = 0
}

// TableName name in bd
func (User) TableName() string {
	return "users"
}
