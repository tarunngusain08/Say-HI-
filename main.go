package main

import (
	"Say-Hi/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	initDB()
	defer DB.Close()
	initHandlers()
	r := gin.New()
	api := r.Group("/api")

	user := api.Group("/user")

	user.POST("/register", handler.User.RegisterHandler.Register)
	user.POST("/verify-email", handler.User.VerifyEmailHandler.VerifyEmail)
	user.POST("/login", handler.User.LoginHandler.Login)

	r.Use(auth.Middleware())
	user.POST("/logout", handler.User.LogoutHandler.Logout)

	notification := api.Group("/notification")
	notification.POST("send-email", handler.Notification.SendEmailHandler.SendEmail)
	r.Run("localhost:8080")
}
