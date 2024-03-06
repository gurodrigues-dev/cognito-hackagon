package main

import (
	"gin/config"
	"gin/internal/controllers"
	"gin/internal/service"
	"gin/repository"
	"log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	repo, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("error creating repository: %s", err.Error())
	}

	redis, err := repository.NewRedisClient()
	if err != nil {
		log.Fatalf("error creating redis connection: %s", err.Error())
	}

	service := service.New(repo, redis)

	controller := controllers.New(service)

	log.Printf("initing service: %s", config.Name)
	controller.Start()

}
