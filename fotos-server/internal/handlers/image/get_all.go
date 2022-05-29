package imageHandler

import (
	"errors"
	"time"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetAllRequestBody struct {
	EventID string `json:"eventId" binding:"required"`
}

type ResponseImageBody struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Src       string    `json:"src"`
	ClickedAt time.Time `json:"clickedAt"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	HasYou    bool      `json:"hasYou"`
}

type GetAllResponseBody struct {
	EventId uint                `json:"eventId"`
	Images  []ResponseImageBody `json:"images"`
}

func GetAll_POST(c *fiber.Ctx) error {
	requestBody := GetAllRequestBody{}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	db := database.GetDB()

	var event models.Event

	if err := db.Preload("Images").First(&event, requestBody.EventID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Event Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	var responseBody GetAllResponseBody

	for _, image := range event.Images {
		// get all presentUsers of image
		if err := db.Preload("PresentUsers").First(&image).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Image Not Found"})
			}
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
		}

		var responseImageBody ResponseImageBody
		responseImageBody.Id = image.ID
		responseImageBody.Name = image.Name
		responseImageBody.Src = image.Url
		responseImageBody.ClickedAt = image.ClickedAt
		responseImageBody.Width = image.Width
		responseImageBody.Height = image.Height
		responseImageBody.HasYou = false

		for _, presentUser := range image.PresentUsers {
			if presentUser.ID == userId {
				responseImageBody.HasYou = true
				break
			}
		}

		responseBody.Images = append(responseBody.Images, responseImageBody)
	}

	responseBody.EventId = event.ID

	return c.Status(200).JSON(responseBody)

}
