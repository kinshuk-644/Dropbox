package connection

import (
	"fmt"
	"log"

	"github.com/kinshuk-644/Dropbox/backend/config"
	"github.com/kinshuk-644/Dropbox/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DBName,
		config.AppConfig.Database.Port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("🚀 Successfully connected to the database!")

	fmt.Println("Running database migrations...")
	DB.AutoMigrate(&models.FileMetadata{})
}
