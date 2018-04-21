package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name  string
	Jobs  []Job
	Users []User `gorm:"many2many:company_users"`
}
