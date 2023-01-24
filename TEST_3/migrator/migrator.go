package main

import (
	"fmt"
	"log"
	"test_3/TEST_3/initializers"
	"test_3/TEST_3/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Error : Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Publisher{}, &models.User{})
	fmt.Println("Success : Migration complete")
}
