package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HarshalSase/url-shortner/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//loading values from env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	//creating gin framework router
	router := gin.Default()

	//setting up routes using router
	setupRouters(router)

	//reading port from env file
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
}
