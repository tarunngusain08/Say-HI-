// Import necessary packages
package main

import (
	"Say-Hi/config"
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
	user.POST("register", config.handler.User.RegisterHandler.Register)
	user.POST("verify-email", config.handler.User.VerifyEmailHandler.VerifyEmail)

	notification := api.Group("/notification")
	notification.POST("send-email", config.handler.Notification.SendEmailHandler.SendEmail)
	r.Run("localhost:8080")
}
