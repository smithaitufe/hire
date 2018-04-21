package context

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/smithaitufe/hire/hire/models"
)

func OpenDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(config.DB.Engine, fmt.Sprintf("host=%s port=%s dbname=%s username=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.Name, config.DB.Username, config.DB.Password))
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Job{},
		&models.Vacancy{},
		&models.JobApplication{},
	)
}
