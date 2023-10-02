package routes

import (
	"github.com/go-redis/redis/v8"                      // Importing the Redis client library
	"github.com/gofiber/fiber/v2"                       // Importing the Fiber web framework
	"github.com/theritikchoure/go-urlshortner/database" // Importing your database package
)

// ResolveURL resolves the shortened URL and redirects to the original URL.
func ResolveURL(c *fiber.Ctx) error {
	// Get the "url" parameter from the route
	url := c.Params("url")

	// Create a Redis client for read operations (index 0)
	r := database.CreateClient(0)
	defer r.Close() // Close the Redis client connection when the function exits

	// Retrieve the original URL associated with the short URL from Redis
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		// Handle the case where the short URL is not found in Redis
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short URL not found",
		})
	} else if err != nil {
		// Handle any other errors that may occur when connecting to the database
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to the database",
		})
	}

	// Create a Redis client for incrementing a counter (index 1)
	rInr := database.CreateClient(1)
	defer rInr.Close() // Close the Redis client connection when the function exits

	// Increment a counter (e.g., for tracking usage statistics)
	_ = rInr.Incr(database.Ctx, "counter")

	// Redirect the client to the original URL with a 301 status code (Moved Permanently)
	return c.Redirect(value, 301)
}
