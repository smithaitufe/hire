package models

import "github.com/jinzhu/gorm"

type JobApplication struct {
	gorm.Model
	JobID       uint
	Job         Job
	UserID      uint
	User        User
	ReferenceNo string
	Invited     bool
}
