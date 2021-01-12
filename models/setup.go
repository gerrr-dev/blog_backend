// models/setup.go

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB
func ConnectDataBase() {
//Loading environment
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s", dbHost, username, dbName, password) //Build connection string
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	DB = conn
	
//Database migration
	DB.Debug().AutoMigrate(&Article{})
	log.Println("Database successfully connected")
}