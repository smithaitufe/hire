package context

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/smithaitufe/hire/hire/models"
)

func OpenDB(config *Config) *gorm.DB {
	db, err := gorm.Open(config.DB.Engine, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.Name, config.DB.User, config.DB.Password))
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	return db
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
