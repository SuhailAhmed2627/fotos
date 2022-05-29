package imageRoutes

import (
	imageHandler "github.com/SuhailAhmed2627/fotos-server/internal/handlers/image"
	"github.com/SuhailAhmed2627/fotos-server/internal/middlewares/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupImageRoutes(router fiber.Router) {
	imageRouter := router.Group("/image")

	imageRouter.Post("/upload", auth.AuthCheck, imageHandler.Upload_POST)

	imageRouter.Post("/get_all", auth.AuthCheck, imageHandler.GetAll_POST)

	imageRouter.Post("/get", auth.AuthCheck, imageHandler.Get_POST)
}
