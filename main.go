package main

import (
	"log"
	"os"

	"github.com/sonu31/expreimnet-go-lang-with-mongoDb/routes"

	"github.com/sonu31/expreimnet-go-lang-with-mongoDb/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()

	routes.UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running at http://localhost:" + port)
	router.Run(":" + port)
}
