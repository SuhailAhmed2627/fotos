package userHandler

import (
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/SuhailAhmed2627/fotos-server/database"
	"github.com/SuhailAhmed2627/fotos-server/internal/models"
	"github.com/SuhailAhmed2627/fotos-server/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func validatePassword(password string) bool {
	var (
		hasMinLen = false
		hasNumber = false
	)
	if len(password) >= 6 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasNumber
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func Signup_POST(c *fiber.Ctx) error {

	var signupRequest SignupRequest

	if err := c.BodyParser(&signupRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Request"})
	}

	if !validateEmail(signupRequest.Email) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Email address is invalid"})
	}

	if !validatePassword(signupRequest.Password) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Password is Invalid"})
	}

	userEmail := strings.Split(signupRequest.Email, "@")
	userEmailBeforeAt := userEmail[0]
	userEmailAfterAt := userEmail[1]
	userEmailBeforePlus := strings.Split(userEmailBeforeAt, "+")[0]
	regExStr := `[^A-Za-z0-9]`
	re := regexp.MustCompile(regExStr)
	userEmailStripped := re.ReplaceAllString(userEmailBeforePlus, "")
	finalEmail := userEmailStripped + "@" + userEmailAfterAt

	db := database.GetDB()
	emails := []models.UserRegistration{}

	dbResponse := db.Where("email= ?", finalEmail).Find(&emails)

	if len(emails) > 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Email Already Exist"})
	}

	if dbResponse.Error != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Internal Error. Try Again"})
	}

	username := []models.UserRegistration{}
	dbResponse = db.Where("username= ?", signupRequest.Username).Find(&username)

	if len(username) > 0 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Username Already Exist"})
	}

	if dbResponse.Error != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Username Already Exist"})
	}

	var isEmailVerified bool = false
	if os.Getenv("APP_ENV") != "develepment" {
		isEmailVerified = true
	}

	userRegistration := models.UserRegistration{
		Username:        signupRequest.Username,
		Name:            signupRequest.Name,
		Email:           signupRequest.Email,
		Password:        utils.HashPwd(signupRequest.Password),
		EmailToken:      utils.HashPwd(signupRequest.Email),
		IsEmailVerified: isEmailVerified,
	}

	dbResponse = db.Create(&userRegistration)

	if dbResponse.Error != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error in user signup!"})
	}

	// if os.Getenv("APP_ENV") != "develepment" {
	// 	emailSent := utils.SendMail(utils.GetVerificationEmailMessage(signupRequest.Email, userRegistration.EmailToken))

	// 	if emailSent != 1 {
	// 		dbResponse = db.Delete(&userRegistration)

	// 		if dbResponse.Error != nil {
	// 			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Some error occured!"})
	// 		}

	// 		return c.Status(418).JSON(fiber.Map{"status": "error", "message": "User not registered, Email couldn't be sent"})
	// 	}
	// }

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Registered"})
}
