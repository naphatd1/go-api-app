package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/naphat/gobapi/cmd/api/internal/adapters/db"
	"github.com/naphat/gobapi/config"
)

func main() {
// Load .env config
	cfg := config.LoadConfig()

	// Create a new Fiber instance
	app := fiber.New()

	// Connect DB & Auto Migrate Models
	db.Connect(cfg)

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from Fiber Next Commerce!"})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
	log.Println("✅ Server is running on port", cfg.AppPort)
}
