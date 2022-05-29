package main

import (
	"fmt"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	server := fiber.New()

	server.Use(cors.New())

	database.ConnectDB()

	models.MigrateDB()

	router.SetupRoutes(server)

	// Listen on PORT 3000
	server.Listen(":3000")
}
