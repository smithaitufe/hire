package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName       string `gorm:"size:50"`
	LastName        string `gorm:"size:50"`
	Password        string
	JobApplications []JobApplication
	Email           string
	PhoneNumber     string    `gorm:"size:30"`
	DialingCode     string    `gorm:"size:20"`
	Country         string    `gorm:"size:100"`
	Companies       []Company `gorm:"many2many:company_users"`
}
