package db

import (
	"fmt"
	"os"

	"github.com/hngphanminh147/gin_api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetupDatabaseConnection() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env config")
	}

	dbUser := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPw,
		dbHost,
		dbPort,
		dbName)
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Tag{})

	Db = db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, errDB := db.DB()

	if errDB != nil {
		panic("Failed to close connection")
	}

	dbSql.Close()
}
