package eventHandler

import (
	"errors"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetAllResponseEvent struct {
	Id          uint   `json:"eventId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
}

type GetAllResponseBody struct {
	Events []GetAllResponseEvent `json:"events"`
}

func GetAll_GET(c *fiber.Ctx) error {
	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	db := database.GetDB()

	var user models.User
	if err := db.Preload("Events").First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User Not Found"})
		}
	}

	if len(user.Events) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "No Events Found"})
	}

	var creator models.User

	for _, event := range user.Events {
		if err := db.First(&creator, event.CreatorID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
		}
		event.Creator = creator
	}

	var responseBody GetAllResponseBody

	for _, event := range user.Events {
		var responseEvent GetAllResponseEvent
		responseEvent.Id = event.ID
		responseEvent.Name = event.Name
		responseEvent.Description = event.Description
		responseEvent.CreatedBy = event.Creator.Name
		responseBody.Events = append(responseBody.Events, responseEvent)
	}

	return c.Status(200).JSON(responseBody)
}
