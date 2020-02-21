package helpers

import (
	"fmt"
	"regexp"
	"time"

	"github.com/asaskevich/govalidator"
)

var namePREG, _ = regexp.Compile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")

func nameValidator(i interface{}, o interface{}) bool {
	v := fmt.Sprintf("%v", i)
	return namePREG.MatchString(v)
}

func dateValidator(i interface{}, o interface{}) bool {
	switch i.(type) {
	case time.Time:
		v := i.(time.Time)
		n := time.Now()
		if v.Year() < n.Year()-100 {
			return false
		}
		if v.Year() > n.Year()-10 {
			return false
		}
		return true
	}
	return false
}

// Init init the helpers
func Init() {
	nameValidator := govalidator.CustomTypeValidator(nameValidator)
	govalidator.CustomTypeTagMap.Set("name", nameValidator)
	govalidator.CustomTypeTagMap.Set("date", dateValidator)
}
