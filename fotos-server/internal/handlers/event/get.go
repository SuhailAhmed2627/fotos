package eventHandler

import (
	"errors"
	"strconv"
	"time"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetRequestBody struct {
	EventID string `json:"eventId" binding:"required"`
}

type ImageResponseBody struct {
	Id        uint      `json:"id"`
	Url       string    `json:"url"`
	ClickedAt time.Time `json:"clickedAt"`
}

type ResponseEvent struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedBy   string              `json:"createdBy"`
	Url         string              `json:"url"`
	Images      []ImageResponseBody `json:"images"`
}

type GetResponseBody struct {
	Event     ResponseEvent `json:"event"`
	IsCreator bool          `json:"isCreator"`
}

func Get_POST(c *fiber.Ctx) error {
	requestBody := GetRequestBody{}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}
	eventId, err := strconv.ParseUint(requestBody.EventID, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

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
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Getting User"})
	}

	var userIsPartOfEvent bool

	if len(user.Events) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User is not part of any Event"})
	}

	for _, userEvent := range user.Events {

		if userEvent.ID == uint(eventId) {
			userIsPartOfEvent = true
			break
		}
	}

	if !userIsPartOfEvent {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User Not Part of Event"})
	}

	var event models.Event
	if err := db.Preload("Images").Preload("Creator").First(&event, requestBody.EventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Event Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Getting Event"})
	}

	var responseBody GetResponseBody
	responseBody.Event.Name = event.Name
	responseBody.Event.Description = event.Description
	responseBody.Event.CreatedBy = event.Creator.Name
	responseBody.Event.Url = event.Url
	responseBody.IsCreator = event.Creator.ID == userId

	for _, image := range event.Images {
		responseBody.Event.Images = append(responseBody.Event.Images, ImageResponseBody{
			Id:        image.ID,
			Url:       image.Url,
			ClickedAt: image.ClickedAt,
		})
	}

	return c.JSON(responseBody)
}
