package imageHandler

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
	ImageId string `json:"imageId" binding:"required"`
}

type GetResponseBody struct {
	Name          string    `json:"name"`
	Src           string    `json:"src"`
	ClickedAt     time.Time `json:"clickedAt"`
	Width         int       `json:"width"`
	Height        int       `json:"height"`
	FaceProcessed bool      `json:"faceProcessed"`
	Uploadername  string    `json:"uploaderName"`
	HasYou        bool      `json:"hasYou"`
}

func Get_POST(c *fiber.Ctx) error {
	requestBody := GetRequestBody{}

	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	db := database.GetDB()

	temp, err := strconv.ParseUint(requestBody.ImageId, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}
	var image = models.Image{
		ID: uint(temp),
	}

	if err := db.Preload("PresentUsers").Preload("Uploader").First(&image).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Image Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	var responseBody GetResponseBody
	responseBody.Name = image.Name
	responseBody.Src = image.Url
	responseBody.ClickedAt = image.ClickedAt
	responseBody.Width = image.Width
	responseBody.Height = image.Height
	responseBody.FaceProcessed = image.FaceProcessed
	responseBody.Uploadername = image.Uploader.Name
	responseBody.HasYou = false

	for _, user := range image.PresentUsers {
		if user.ID == userId {
			responseBody.HasYou = true
			break
		}
	}

	return c.Status(200).JSON(responseBody)
}

func Uint(u uint64, err error) {
	panic("unimplemented")
}
