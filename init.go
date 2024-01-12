package main

import (
	"Say-Hi/auth"
	"Say-Hi/config"
	NotificationHandler "Say-Hi/notification/handler"
	NotificationService "Say-Hi/notification/service"
	"Say-Hi/user/external"
	UserHandler "Say-Hi/user/handler"
	UserRepo "Say-Hi/user/repo"
	UserService "Say-Hi/user/service"
	"database/sql"
	"fmt"
	"log"
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
)

type Handler struct {
	User
	Notification
	Message
	ChatHistory
}

type User struct {
	RegisterHandler       *UserHandler.RegisterHandler
	VerifyEmailHandler    *UserHandler.VerifyEmailHandler
	LoginHandler          *UserHandler.LoginHandler
	LogoutHandler         *UserHandler.LogoutHandler
	ForgotPasswordHandler *UserHandler.ForgotPasswordHandler
}

type Notification struct {
	SendEmailHandler *NotificationHandler.SendEmailHandler
}

type Message struct {
}

type ChatHistory struct {
}

func initHandlers() {

	config.Init()
	registerRepo := UserRepo.NewRegisterRepo(DB)
	verifyEmailRepo := UserRepo.NewVerifyEmailRepo(DB)
	loginRepo := UserRepo.NewLoginRepo(DB)
	logoutRepo := UserRepo.NewLogoutRepo(DB)
	forgotPasswordRepo := UserRepo.NewForgotPasswordRepo(DB)

	registerService := UserService.NewRegisterService(registerRepo)
	loginService := UserService.NewLoginService(loginRepo)
	logoutService := UserService.NewLogoutService(logoutRepo)
	verifyEmailService := UserService.NewVerifyEmailService(verifyEmailRepo)
	forgotPasswordService := UserService.NewForgotPasswordService(forgotPasswordRepo)
	emailService := external.NewEmailService(config.Config.MaxRetries, time.Duration(config.Config.BaseDelay), time.Duration(config.Config.MaxDelay))
	sendEmailService := NotificationService.NewSendEmailService()

	registerHandler := UserHandler.NewRegisterHandler(registerService, emailService)
	jwtHandler := auth.NewJWT()
	loginHandler := UserHandler.NewLoginHandler(loginService, jwtHandler)
	logoutHandler := UserHandler.NewLogoutHandler(logoutService)
	verifyEmailHandler := UserHandler.NewVerifyEmailHandler(verifyEmailService)
	forgotPasswordHandler := UserHandler.NewForgotPasswordHandler(forgotPasswordService, emailService)
	notificationHandler := NotificationHandler.NewSendEmailHandler(sendEmailService)

	handler = new(Handler)

	handler.User.RegisterHandler = registerHandler
	handler.User.VerifyEmailHandler = verifyEmailHandler
	handler.User.LoginHandler = loginHandler
	handler.User.LogoutHandler = logoutHandler
	handler.User.ForgotPasswordHandler = forgotPasswordHandler

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
