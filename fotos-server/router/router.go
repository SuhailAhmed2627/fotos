package router

import (
	eventRoutes "github.com/SuhailAhmed2627/fotos-server/internal/routes/event"
	imageRoutes "github.com/SuhailAhmed2627/fotos-server/internal/routes/image"
	userRoutes "github.com/SuhailAhmed2627/fotos-server/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRoutes.SetupUserRoutes(api)
	eventRoutes.SetupEventRoutes(api)
	imageRoutes.SetupImageRoutes(api)
}
