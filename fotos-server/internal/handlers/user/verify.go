package userHandler

import (
	"time"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
)

type EmailVerification struct {
	EmailToken string `json:"emailToken" binding:"required"`
}

func Verify_POST(c *fiber.Ctx) error {
	requestBody := EmailVerification{}
	var status = true
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}
	db := database.GetDB()
	dbResponse := db.Where("email_token= ?", requestBody.EmailToken).Find(&models.UserRegistration{}).Updates(&models.UserRegistration{IsEmailVerified: status, UpdatedAt: time.Now()})
	if dbResponse.Error != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Wrong email token"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Verified Successfully"})
}
