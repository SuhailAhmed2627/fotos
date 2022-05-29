package eventHandler

import (
	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateRequestBody struct {
	EventId     uint   `json:"eventId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Update_POST(c *fiber.Ctx) error {
	requestBody := UpdateRequestBody{}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	db := database.GetDB()
	event := models.Event{}

	db.First(&event, requestBody.EventId)

	if err := db.First(&event, requestBody.EventId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Event Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Updating Event"})
	}

	if event.CreatorID != userId {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Only Creator can Edit Details"})
	}

	event.Description = requestBody.Description
	event.Name = requestBody.Name

	db.Save(&event)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Details Updated"})
}
