package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

// User model of users
type User struct {
	ID   uint      `json:"id"   xorm:"id bigint not null autoincr pk"`
	Name string    `json:"name" valid:"required,name" xorm:"name varchar(100) not null"`                      // ○ Nombre Completo (Requerido)
	DNI  string    `json:"dni"  valid:"required,numeric,length(8|12)" xorm:"dni varchar(20) not null unique"` // ○ DNI o cédula (Requerido)
	Date time.Time `json:"date" valid:"required,date"`                                                        // ○ Fecha de Nacimiento. (El usuario puede o no enviar esta información)
}

type user struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	DNI  string `json:"dni"`
	Date string `json:"date"`
}

func (that *User) UnmarshalJSON(b []byte) error {
	var err error
	m := &user{}
	json.Unmarshal(b, m)
	that.ID = m.ID
	that.Name = m.Name
	that.DNI = m.DNI

	if strings.Contains(m.Date, "T") {
		that.Date, err = time.Parse(time.RFC3339, m.Date)
		that.Date = that.Date.UTC()

	} else {
		that.Date, err = time.Parse(time.RFC3339, m.Date+"T05:00:00Z")
		that.Date = that.Date.UTC()
	}
	if err != nil {
		return errors.New("Date is not valid")
	}
	return nil
}

func (that User) MarshalJSON() ([]byte, error) {
	m := &user{}
	m.ID = that.ID
	m.Name = that.Name
	m.DNI = that.DNI
	m.Date = that.Date.Format(time.RFC3339)[:10]
	return json.Marshal(m)
}

// CleanForInsert clean for insert
func (that *User) CleanForInsert() {
	that.ID = 0
}

// Validate validate model
func (that User) Validate() error {
	_, err := govalidator.ValidateStruct(that)
	return err
}

// TableName name in bd
func (User) TableName() string {
	return "users"
}
