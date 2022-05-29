package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type Config struct {
	Development ConfigDetails
	Test        ConfigDetails
	Production  ConfigDetails
}

type ConfigDetails struct {
	Username           string
	Password           string
	Database           string
	Host               string
	Port               string
	Dialect            string
	BaseUrl            string
	AwsBucket          string
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
}

func ReadConfig() (string, ConfigDetails) {

	env := os.Getenv("APP_ENV")
	var config Config
	var currentConfig ConfigDetails

	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(color.RedString("Error reading file - config.json"), err)
		return "fail", currentConfig
	}

	fmt.Println(color.GreenString("Reading file - config.json......."))

	json.Unmarshal([]byte(data), &config)

	if env == "production" {
		currentConfig = config.Production
	} else {
		currentConfig = config.Development
	}
	return "success", currentConfig
}
