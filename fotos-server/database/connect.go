package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

type Config struct {
	Development ConfigDetails
	Production  ConfigDetails
}

type ConfigDetails struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Dialect  string
}

// ConnectDB connect to db
func ConnectDB() {
	env := os.Getenv("APP_ENV")
	var config Config
	var currentConfig ConfigDetails

	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(color.RedString("Error reading file - config.json"), err)
		return
	}

	fmt.Println(color.GreenString("Reading file - config.json......."))

	json.Unmarshal([]byte(data), &config)

	if env == "production" {
		currentConfig = config.Production
	} else {
		currentConfig = config.Development
	}

	var connectionString string = currentConfig.Username+":"+
	currentConfig.Password+"@tcp("+
	currentConfig.Host+":"+
	currentConfig.Port+")/"+
	currentConfig.Database+"?charset=utf8&parseTime=True"

	db, err = gorm.Open(mysql.Open(connectionString), )

	if err != nil {
		fmt.Println(color.RedString("MySQL connection Failed to Open "), err)
	} else {
		fmt.Println(color.GreenString("MySQL connection Established"))
	}
}