package eventHandler

import (
	"context"
	"errors"
	"strconv"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/internal/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type JoinRequestBody struct {
	Url string `json:"url" binding:"required"`
}

type JoinResponseBody struct {
	EventId     uint   `json:"id"`
	Message     string `json:"message" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Join_POST(c *fiber.Ctx) error {
	requestBody := JoinRequestBody{}

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

	if err := db.Where("url = ?", requestBody.Url).Preload("Users").First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Event Not Found"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Joining Event"})
	}

	var alreadyPresent bool = false

	for _, user := range event.Users {
		if user.ID == userId {
			alreadyPresent = true
			break
		}
	}

	responseBody := JoinResponseBody{
		Name:        event.Name,
		Description: event.Description,
		Message:     "Already Part of Event",
	}

	if alreadyPresent {
		return c.Status(200).JSON(responseBody)
	}

	newUser := models.User{}
	if err := db.First(&newUser, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Finding User"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Joining Event"})
	}

	event.Users = append(event.Users, &newUser)

	if err := db.Save(&event).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Joining Event"})
	}

	newUser.Events = append(newUser.Events, &event)

	if err := db.Save(&newUser).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Joining Event"})
	}

	if err := db.Model(&models.Image{}).Where("event_id = ?", event.ID).Update("face_processsed", false).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Joining Event"})
	}

	responseBody.Message = "Joined Event"

	res, config := utils.ReadConfig()

	if res != "success" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	rekog := rekognition.New(rekognition.Options{
		Region:      config.AwsRegion,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")),
	})

	faceIndexResponse, err := rekog.IndexFaces(context.TODO(), &rekognition.IndexFacesInput{
		CollectionId: aws.String("eventFaceId-" + strconv.FormatUint(uint64(event.ID), 10)),
		Image: &types.Image{S3Object: &types.S3Object{
			Bucket: aws.String(config.AwsBucket),
			Name:   &newUser.FaceUrl,
		}},
		MaxFaces: aws.Int32(1),
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if len(faceIndexResponse.FaceRecords) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	faceId := faceIndexResponse.FaceRecords[0].Face.FaceId

	newFace := models.Face{
		EventId:   event.ID,
		OwnerId:   userId,
		AWSFaceID: *faceId,
	}

	if err := db.Create(&newFace).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Creating Face"})
	}

	responseBody.EventId = event.ID

	return c.Status(200).JSON(responseBody)
}
