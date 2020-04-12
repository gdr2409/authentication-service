package user

import (
	"github.com/jinzhu/gorm"
)

// Person model stores user's data
type Person struct {
	gorm.Model
	Name     string `gorm:"size:20;column:name;NOT NULL"`
	Age      int    `gorm:"column:age"`
	Email    string `gorm:"UNIQUE;size:30;column:email;NOT NULL"`
	Password string `gorm:"size:30;column:password;NOT NULL"`
}

// TableName for Person is `person`
func (Person) TableName() string {
	return "person"
}
