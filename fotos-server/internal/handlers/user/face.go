package userHandler

import (
	"context"
	"strconv"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/internal/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func Face_POST(c *fiber.Ctx) error {
	idLocal := c.Locals("id")
	userId, ok := idLocal.(uint)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	files, ok := form.File["files"]
	if !ok {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	for _, file := range files {

		// open the file
		fileTemp, err := file.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid file"})
		}

		// get the file content type
		fileContentType, err := utils.GetFileContentType(fileTemp)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid file"})
		}

		// check if the file is an image
		if fileContentType != "image/jpeg" && fileContentType != "image/png" {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid file type"})
		}
	}

	res, config := utils.ReadConfig()

	if res != "success" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	if len(files) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	client := s3.New(s3.Options{
		Region:      config.AwsRegion,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")),
	})

	uploader := manager.NewUploader(client)

	f, err := files[0].Open()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}
	filekey := "users/face/user-" + strconv.FormatUint(uint64(userId), 10) + "/" + files[0].Filename
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &config.AwsBucket,
		Key:    &(filekey),
		Body:   f,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	db := database.GetDB()

	if result != nil {
		db.Model(&models.User{}).Where("id = ?", userId).Update("face_url", filekey)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Face Updated"})
}
