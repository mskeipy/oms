package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Env       string
	AppName   string
	AppApiUrl string

	PostgresURL      string
	PostgresDatabase string

	RestPort string

	Version string

	JwtSecret string

	MailHost        string
	MailPort        int
	MailUsername    string
	MailPassword    string
	MailFromName    string
	MailFromAddress string

	RabbitMqUrl string
}

var Global = Env{}

func InitApp() {
	LoadEnv()
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		log.Fatal("----------Error loading .env file-----------")
	}

	Global.Env = os.Getenv("ENV")
	Global.AppName = os.Getenv("APP_NAME")
	Global.AppApiUrl = os.Getenv("APP_API_URL")

	Global.PostgresURL = os.Getenv("POSTGRES_URL")
	Global.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

	Global.RestPort = os.Getenv("REST_PORT")

	Global.Version = os.Getenv("API_VERSION")

	Global.JwtSecret = os.Getenv("JWT_SECRET")

	Global.MailHost = os.Getenv("MAIL_HOST")
	Global.MailPort, _ = strconv.Atoi(os.Getenv("MAIL_PORT"))
	Global.MailUsername = os.Getenv("MAIL_USERNAME")
	Global.MailPassword = os.Getenv("MAIL_PASSWORD")
	Global.MailFromName = os.Getenv("MAIL_FROM_NAME")
	Global.MailFromAddress = os.Getenv("MAIL_FROM_ADDRESS")
	Global.RabbitMqUrl = os.Getenv("RABBITMQ_URL")
}
