package initializers

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	DB, err = gorm.Open("postgres", "host=%s user=%s dbname=%s password=%s", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}
