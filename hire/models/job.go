package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Title       string
	Description string
	CompanyID   uint
	Company     Company
	Vacancies   []Vacancy
}
