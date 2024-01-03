package main

import (
	NotificationHandler "Say-Hi/notification/handler"
	NotificationService "Say-Hi/notification/service"
	"Say-Hi/user/external"
	UserHandler "Say-Hi/user/handler"
	UserRepo "Say-Hi/user/repo"
	UserService "Say-Hi/user/service"
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

const (
	userName = "tgusain"
	dbName   = "tgusain"
	disable  = "disable"
)

var (
	DB      *sql.DB
	handler *Handler
	Config  *Configuration
)

type Handler struct {
	User
	Notification
	Message
	ChatHistory
}

type User struct {
	RegisterHandler    *UserHandler.RegisterHandler
	VerifyEmailHandler *UserHandler.VerifyEmailHandler
}

type Notification struct {
	SendEmailHandler *NotificationHandler.SendEmailHandler
}

type Message struct {
}

type ChatHistory struct {
}

type Configuration struct {
	MaxRetries int `yaml:"maxRetries"`
	BaseDelay  int `yaml:"baseDelay"`
	MaxDelay   int `yaml:"maxDelay"`
}

func initDefaultValues() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("Error reading YAML file")
	}

	// Parse the YAML content into a Configuration struct
	Config = new(Configuration)
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Error unmarshalling YAML")
	}
}

func initHandlers() {

	initDefaultValues()
	registerRepo := UserRepo.NewRegisterRepo(DB)
	verifyEmailRepo := UserRepo.NewVerifyEmailRepo(DB)

	registerService := UserService.NewRegisterService(registerRepo)
	verifyEmailService := UserService.NewVerifyEmailService(verifyEmailRepo)
	emailService := external.NewEmailService(Config.MaxRetries, time.Duration(Config.BaseDelay), time.Duration(Config.MaxDelay))
	sendEmailService := NotificationService.NewSendEmailService()

	registerHandler := UserHandler.NewRegisterHandler(registerService, emailService)
	verifyEmailHandler := UserHandler.NewVerifyEmailHandler(verifyEmailService)
	notificationHandler := NotificationHandler.NewSendEmailHandler(sendEmailService)

	handler = new(Handler)

	handler.User.RegisterHandler = registerHandler
	handler.User.VerifyEmailHandler = verifyEmailHandler

	handler.Notification.SendEmailHandler = notificationHandler
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", userName, dbName, disable)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}

func initDB() {
	var err error
	DB, err = connectDB()
	if err != nil {
		log.Fatal(err)
	}
}
