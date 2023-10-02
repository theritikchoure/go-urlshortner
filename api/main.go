package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"                     // Importing the Fiber web framework
	"github.com/gofiber/fiber/v2/middleware/logger"   // Importing the Fiber logger middleware
	"github.com/joho/godotenv"                        // Importing the godotenv library for loading environment variables
	"github.com/theritikchoure/go-urlshortner/routes" // Importing your application's route handlers
)

// setupRoutes configures the application's routes and handlers.
func setupRoutes(app *fiber.App) {
	// Define a route that resolves shortened URLs by their custom short ID
	app.Get("/:url", routes.ResolveURL)

	// Define an API route to create shortened URLs
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	// Load environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Create a new Fiber application instance
	app := fiber.New()

	// Use the Fiber logger middleware to log HTTP requests
	app.Use(logger.New())

	// Configure the application's routes by calling the setupRoutes function
	setupRoutes(app)

	// Get the port from the environment variables, default to 3000 if not specified
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// Start the Fiber application and listen on the specified port
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
