package utils

import (
	"context"
	"strconv"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

func FaceDetect(eventId uint) {
	db := database.GetDB()

	var images []models.Image

	if err := db.Where("face_processed = ? AND event_id = ?", false, eventId).Find(&images).Error; err != nil {
		return
	}

	if len(images) == 0 {
		return
	}

	var event models.Event

	if err := db.Preload("Users").First(&event, eventId).Error; err != nil {
		return
	}

	res, config := ReadConfig()

	if res != "success" {
		return
	}

	// Create rekognition manager
	rekog := rekognition.New(rekognition.Options{
		Region:      config.AwsRegion,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")),
	})

	for _, image := range images {
		req := rekognition.SearchFacesByImageInput{
			CollectionId: aws.String("eventFaceId-" + strconv.FormatUint(uint64(event.ID), 10)),
			Image: &types.Image{S3Object: &types.S3Object{
				Bucket: aws.String(config.AwsBucket),
				Name:   &image.Key,
			}},
		}

		// Call rekognition SearchFacesByImage
		res, err := rekog.SearchFacesByImage(context.TODO(), &req)

		if err != nil {
			continue
		}

		var faceMatch bool = true

		if len(res.FaceMatches) == 0 {
			faceMatch = false
		} else if *(res.FaceMatches[0].Face.Confidence) < float32(0.5) {
			faceMatch = false
		}

		if faceMatch {

			userFaceId := res.FaceMatches[0].Face.FaceId

			// Find Face with ownerId and eventId
			var face models.Face

			if err := db.Where("aws_face_id = ? AND event_id = ?", *userFaceId, eventId).Preload("Owner").First(&face).Error; err != nil {
				continue
			}

			image.PresentUsers = append(image.PresentUsers, &face.Owner)
		}

		image.FaceProcessed = true

		if err := db.Save(&image).Error; err != nil {
			continue
		}
	}
}
