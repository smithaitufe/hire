package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Vacancy struct {
	gorm.Model
	JobID            uint
	Job              Job
	StartDate        time.Time
	ClosingDate      time.Time
	ShortDescription string
	LongDescription  string
}
