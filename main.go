package main

import (
	"log"
	"os"
	"os/signal"
	"portfolio_form/config"
	"portfolio_form/controllers"
	"portfolio_form/repo"
	"portfolio_form/services"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitializeEnv()
	database, err := config.InitializeDb(config.Envs.MongoURL)
	if err != nil {

	}

	repo := repo.InitializeFormRepo(database)

	services := services.InitializeFormService(repo)

	handlers := controllers.InitializeFormController(services)

	r := gin.Default()

	// form routes
	r.POST("/createForm", handlers.CreateFormDetails)
	r.GET("/getForm/:email", handlers.GetAllFormDetailsByEmail)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello machi! Gin API working🔥",
			"status":  "success",
		})
	})

	go func() {
		log.Println("✅ Server is running on port 8080")
		if err := r.Run(); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("🔻 Shutting down server...")
	config.DisconnectMongoDB()
	log.Println("✅ Server stopped gracefully")

}
