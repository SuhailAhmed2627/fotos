package imageHandler

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"strconv"
	"sync"
	"time"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/internal/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type DimensionsObject struct {
	Dimensions []Dimensions `json:"dimensions"`
}

type ClickedAtObject struct {
	ClickedAts []time.Time `json:"clickedAts"`
}

func Upload_POST(c *fiber.Ctx) error {

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
	temp, err := strconv.ParseUint(form.Value["eventId"][0], 10, 8)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}
	eventId := uint(temp)
	clickedAtsTemp := form.Value["clickedAts"][0]

	var clickedAtsObject ClickedAtObject
	_ = json.Unmarshal([]byte(clickedAtsTemp), &clickedAtsObject)

	clickedAts := clickedAtsObject.ClickedAts

	dimensionsTemp := form.Value["dimensions"][0]

	var dimensionsObject DimensionsObject
	_ = json.Unmarshal([]byte(dimensionsTemp), &dimensionsObject)

	dimensions := dimensionsObject.Dimensions

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

	client := s3.New(s3.Options{
		Region:      config.AwsRegion,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")),
	})

	uploader := manager.NewUploader(client)

	var waitGroup sync.WaitGroup
	var transcationMutex = sync.Mutex{}
	threadId := -1

	db := database.GetDB()

	var errorsArray []int

	for index, file := range files {
		waitGroup.Add(1)
		threadId++

		go func(file *multipart.FileHeader, threadId int, index int, clickedAts []time.Time, dimensions []Dimensions) {
			defer waitGroup.Done()
			f, err := file.Open()
			if err != nil {
				errorsArray = append(errorsArray, index)
				return
			}
			filekey := "event-" + strconv.FormatUint(uint64(eventId), 10) + "/images/" + uuid.New().String() + file.Filename
			result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
				Bucket: &config.AwsBucket,
				Key:    &(filekey),
				Body:   f,
			})
			if err != nil {
				errorsArray = append(errorsArray, index)
			}
			if result != nil {
				transcationMutex.Lock()
				db.Create(&models.Image{
					EventID:    eventId,
					Url:        result.Location,
					Name:       file.Filename,
					Width:      dimensions[index].Width,
					Height:     dimensions[index].Height,
					UploaderID: userId,
					ClickedAt:  clickedAts[index],
					Key:        filekey,
				})
				transcationMutex.Unlock()
			}
		}(file, threadId, index, clickedAts, dimensions)
	}
	waitGroup.Wait()

	go utils.FaceDetect(eventId)

	if len(errorsArray) > 0 {
		return c.Status(300).JSON(fiber.Map{"status": "error", "message": "Error Uploading Files", "errors": errorsArray})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "file uploaded successfully"})

}
