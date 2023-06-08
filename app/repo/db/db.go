package db

import (
	"fmt"
	"log"
	"os"
	CarInfoEntity "project-B/app/entity"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// declare a Db object, where we can use throughout the entity package
// so in blog.go, we have access to this object
var connInfo connection

// a struct to hold all the Db connection information
type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDB() *gorm.DB {
	connection, err := gorm.Open("postgres", getPsqlInfo(connInfo))
	if err != nil {
		log.Fatal("Could not connect to database : %s", err)
		return nil
	}

	return connection
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	// Retrieving our connection information from our environment variables
	connInfo = connection{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PWD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	// try to open our postgresql connection with our connection info
	db, err := gorm.Open("postgres", getPsqlInfo(connInfo))
	if err != nil {
		log.Fatal("Could not connect to database : %s", err)
	} else {
		fmt.Printf("DB is open\n")
	}

	//Close connection when main function ends
	db.AutoMigrate(&CarInfoEntity.CarInfoEntity{}) //Database migration

}

// Take our connection struct and convert to a string for our Db connection info
func getPsqlInfo(info connection) string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", info.Host, info.Port, info.User, info.Password, info.DBName)

}
