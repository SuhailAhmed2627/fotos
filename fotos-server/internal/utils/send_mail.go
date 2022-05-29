package utils

import (
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(message *mail.SGMailV3) int {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return 0
	}

	if response.StatusCode != 202 {
		log.Println(err)
		return 0
	}
	return 1
}

func GetVerificationEmailMessage(userEmail, emailToken string) *mail.SGMailV3 {
	from := mail.NewEmail("FOTOS", "suhailahmed2001sam@gmail.com")
	subject := "Fotos Email Verification"
	to := mail.NewEmail(userEmail, userEmail)
	plainTextContent := "Welcome to Fotos. To complete your Fotos, verify your Email"
	verifLink := "http://localhost:3000/api/user/verify?token=" + emailToken
	status, config := ReadConfig()
	if status == "success" {
		verifLink = config.BaseUrl + "/verify?token=" + emailToken
	}
	htmlContent := "<strong>" + verifLink + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	return message
}
