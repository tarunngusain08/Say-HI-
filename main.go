// Import necessary packages
package main

import (
	"Say-Hi/types"
	userHandler "Say-Hi/user/handler"
	userRepo "Say-Hi/user/repo"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB      *sql.DB
	Handler *types.Handler
)

func connectDB() (*sql.DB, error) {
	connStr := "user=tgusain dbname=tgusain sslmode=disable"
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

func initHandlers() {
	userRepo := userRepo.NewRegisterRepo(DB)
	userHandler := userHandler.NewRegisterHandler(userRepo)
	Handler = new(types.Handler)
	Handler.User.RegisterHandler = userHandler
}

func main() {
	initDB()
	initHandlers()
	defer DB.Close()
	r := gin.New()
	api := r.Group("/api")
	api.POST("register", Handler.RegisterHandler.Register)
	r.Run("localhost:8080")
}
