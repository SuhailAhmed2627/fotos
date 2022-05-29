package eventRoutes

import (
	eventHandler "github.com/SuhailAhmed2627/fotos-server/internal/handlers/event"
	"github.com/SuhailAhmed2627/fotos-server/internal/middlewares/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(router fiber.Router) {
	eventRouter := router.Group("/event")

	eventRouter.Post("/create", auth.AuthCheck, eventHandler.Create_POST)

	eventRouter.Post("/update", auth.AuthCheck, eventHandler.Update_POST)

	eventRouter.Post("/join", auth.AuthCheck, eventHandler.Join_POST)

	eventRouter.Post("/get", auth.AuthCheck, eventHandler.Get_POST)

	eventRouter.Get("/get_all", auth.AuthCheck, eventHandler.GetAll_GET)

	eventRouter.Post("/get_users", auth.AuthCheck, eventHandler.GetUsers_POST)
}
