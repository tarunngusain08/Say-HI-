// Import necessary packages
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	initDB()
	initHandlers()
	defer DB.Close()
	r := gin.New()
	api := r.Group("/api")

	user := api.Group("/user")
	user.POST("register", handler.User.RegisterHandler.Register)
	user.POST("verify-email", handler.User.VerifyEmailHandler.VerifyEmail)

	notification := api.Group("/notification")
	notification.POST("send-email", handler.Notification.SendEmailHandler.SendEmail)
	r.Run("localhost:8080")
}
