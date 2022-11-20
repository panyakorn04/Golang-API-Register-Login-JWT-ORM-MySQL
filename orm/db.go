package orm

import (
	"log"
	"os"
	"time"

	"example.com/greetings/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a global variable that represents the database connection
var DB *gorm.DB

// ConnectDatabase is a function that connects to the database
func ConnectDatabase() {
	var err error
	dsn := os.Getenv("MYSQL_DNS")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

 	// Migrate the schema
	DB.AutoMigrate(&model.User{})
}

