package userRoutes

import (
	userHandler "github.com/SuhailAhmed2627/fotos-server/internal/handlers/user"
	"github.com/SuhailAhmed2627/fotos-server/internal/middlewares/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	userRouter := router.Group("/user")

	userRouter.Post("/login", userHandler.Login_POST)

	userRouter.Post("/verify", userHandler.Verify_POST)

	userRouter.Post("/signup", userHandler.Signup_POST)

	userRouter.Post("/face", auth.AuthCheck, userHandler.Face_POST)
}
