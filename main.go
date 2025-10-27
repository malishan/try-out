package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/try-out/api"
	"github.com/try-out/external/database"
	"github.com/try-out/external/mail"
	"github.com/try-out/service"
)

func init() {
	initDatabase()
	initMail()
}

func main() {
	gin.SetMode(gin.DebugMode)

	db := database.GetClient()
	mail := mail.GetClient()

	service.InitService(db, mail)

	go startRouter()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}

func initDatabase() {
	database.InitializeDatabase()
}

func initMail() {
	mail.InitializeMail()
}

func startRouter() {
	router := api.GetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Router initialization failed", err)
	}
}
