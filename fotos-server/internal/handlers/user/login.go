package userHandler

import (
	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/middlewares/auth"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/internal/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Event struct {
	EventID uint   `json:"eventId"`
	Name    string `json:"name"`
}

type LoginResponseBody struct {
	Username   string  `json:"username"`
	Name       string  `json:"name"`
	Events     []Event `json:"events"`
	UserToken  string  `json:"userToken"`
	FirstLogin bool    `json:"firstLogin"`
}

func Login_POST(c *fiber.Ctx) error {
	var requestBody LoginRequestBody

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	db := database.GetDB()

	var user models.User
	var userRegistration models.UserRegistration
	var isNewUser bool = false
	dbResponse := db.Where("email = ?", requestBody.Email).Preload("Events").First(&user)

	if dbResponse.Error != nil && dbResponse.Error != gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error, Try Again"})
	}

	dbResponse2 := db.Where("email = ?", requestBody.Email).First(&userRegistration)

	if dbResponse.Error != nil && dbResponse.Error == gorm.ErrRecordNotFound {

		if dbResponse2.Error != nil && dbResponse2.Error != gorm.ErrRecordNotFound {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error, Try Again"})
		}
		if dbResponse2.Error != nil && dbResponse2.Error == gorm.ErrRecordNotFound {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Credentials"})
		}
		if !userRegistration.IsEmailVerified {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Email not Verified"})
		}

		isNewUser = true
	}

	if utils.HashPwd(requestBody.Password) == userRegistration.Password {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Credentials"})
	}

	if isNewUser {
		var newUser models.User
		dbResponse := db.Create(&models.User{
			Email:    userRegistration.Email,
			Username: userRegistration.Username,
			Name:     userRegistration.Name,
		}).Scan(&newUser)

		if dbResponse.Error != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
		}
		user = newUser
	}
	tokenString := auth.GenerateToken(user.ID)
	err := db.Model(&user).Updates(&models.User{
		UserToken: tokenString,
	}).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error"})
	}

	var events []Event
	if !isNewUser {
		for i, event := range user.Events {
			events = append(events, Event{
				EventID: event.ID,
				Name:    event.Name,
			})
			if i == 2 {
				break
			}
		}
	} else {
		events = []Event{}
	}

	responseBody := LoginResponseBody{
		Username:   user.Username,
		Name:       user.Name,
		Events:     events,
		UserToken:  tokenString,
		FirstLogin: user.FaceUrl == "NEW_USER",
	}

	return c.Status(200).JSON(responseBody)
}
