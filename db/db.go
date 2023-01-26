package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/yourprojectname/models"
)

var db *gorm.DB

// Connect connects to the database
func Connect() {
	var err error
	dbURL := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable password=" + os.Getenv("DB_PASSWORD")

	db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}, &models.Task{})
}

// Close closes the database connection
func Close() {
	db.Close()
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
