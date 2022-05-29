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

type CreateRequestBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateResponseBody struct {
	EventId     uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Create_POST(c *fiber.Ctx) error {
	requestBody := CreateRequestBody{}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if requestBody.Description == "" || requestBody.Name == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	newEvent := models.Event{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		CreatorID:   userId,
	}

	db := database.GetDB()
	creator := models.User{}

	if err := db.First(&creator, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Unable to find User"})
		}
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Updating Event"})
	}

	newEvent.Users = append(newEvent.Users, &creator)

	if err := db.Create(&newEvent).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Creating Event"})
	}

	// append user to events and save
	creator.Events = append(creator.Events, &newEvent)

	if err := db.Save(&creator).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Updating User"})
	}
	tx := db.Begin()

	res, config := utils.ReadConfig()

	if res != "success" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	rekog := rekognition.New(rekognition.Options{
		Region:      config.AwsRegion,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")),
	})

	_, err := rekog.CreateCollection(context.TODO(), &rekognition.CreateCollectionInput{
		CollectionId: aws.String("eventFaceId-" + strconv.FormatUint(uint64(newEvent.ID), 10)),
	})

	if err != nil {
		_, err := rekog.DeleteCollection(context.TODO(), &rekognition.DeleteCollectionInput{
			CollectionId: aws.String("eventFaceId-" + strconv.FormatUint(uint64(newEvent.ID), 10)),
		})
		if err != nil {
			tx.Rollback()
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
		}
	}

	faceIndexResponse, err := rekog.IndexFaces(context.TODO(), &rekognition.IndexFacesInput{
		CollectionId: aws.String("eventFaceId-" + strconv.FormatUint(uint64(newEvent.ID), 10)),
		Image: &types.Image{S3Object: &types.S3Object{
			Bucket: aws.String(config.AwsBucket),
			Name:   &creator.FaceUrl,
		}},
		MaxFaces: aws.Int32(1),
	})

	if err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if len(faceIndexResponse.FaceRecords) == 0 {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	faceId := faceIndexResponse.FaceRecords[0].Face.FaceId

	// Create a Face with eventId and userId
	newFace := models.Face{
		EventId:   newEvent.ID,
		OwnerId:   userId,
		AWSFaceID: *faceId,
	}

	if err := tx.Create(&newFace).Error; err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error Creating Face"})
	}

	tx.Commit()

	responseBody := CreateResponseBody{
		EventId:     newEvent.ID,
		Name:        newEvent.Name,
		Description: newEvent.Description,
	}

	return c.Status(200).JSON(responseBody)
}
