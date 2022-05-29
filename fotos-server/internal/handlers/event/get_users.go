package eventHandler

import (
	"errors"
	"strconv"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetUserResponseParticipantBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type GetUsersResponseBody struct {
	Participants []GetUserResponseParticipantBody `json:"participants"`
}

type GetUsersRequestBody struct {
	EventID string `json:"eventId"`
}

func GetUsers_POST(c *fiber.Ctx) error {
	requestBody := GetUsersRequestBody{}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if requestBody.EventID == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	eventID, err := strconv.ParseUint(requestBody.EventID, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	db := database.GetDB()

	var event models.Event
	if err := db.Preload("Users").First(&event, eventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Event Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if event.CreatorID != userId {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "You are not the Creator of this Event"})
	}

	if event.Users == nil || len(event.Users) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "No Participants Found"})
	}

	var usersResponseBody GetUsersResponseBody

	for _, user := range event.Users {
		var userResponseBody GetUserResponseParticipantBody
		userResponseBody.Name = user.Name
		userResponseBody.Username = user.Username
		usersResponseBody.Participants = append(usersResponseBody.Participants, userResponseBody)
	}

	return c.Status(200).JSON(usersResponseBody)
}
