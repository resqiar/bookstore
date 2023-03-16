package main

import (
	"fmt"
	"log"
	"os"

	"bookstore/database"
	"bookstore/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load variables from .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Init fiber instance
	server := fiber.New()

	// Connect to PostgreSQL
	database.Connect()

	v1 := server.Group("/v1")

	// Init Routes
	routes.UsersRoutes(v1.(*fiber.Group))
	routes.AuthRoutes(v1.(*fiber.Group))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}

	err := server.Listen(":" + PORT)
	if err != nil {
		fmt.Printf("Problem starting server: %v", err)
	}
}
