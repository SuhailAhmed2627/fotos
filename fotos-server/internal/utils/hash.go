package utils

import "golang.org/x/crypto/bcrypt"

func HashPwd(password string) string {
	pwd := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "error"
	}
	return string(hashedPassword)
}
