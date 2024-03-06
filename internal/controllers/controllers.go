package controllers

import (
	"fmt"
	"gin/config"
	"gin/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service.Service
}

func New(s *service.Service) *controller {
	return &controller{
		service: s,
	}
}

func (ct *controller) getUser(c *gin.Context) {

	user, found := ct.service.ReadUser(c)

	if found {
		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	}

	newUser := ct.service.GenerateRandomLogin()

	err := ct.service.CreateUser(c, &newUser)

	if err != nil {
		log.Printf("error while creating user in redis: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	err = ct.service.SaveUser(c, &newUser)

	if err != nil {
		log.Printf("error while saving user in db: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": &newUser})

}

func (ct *controller) Start() {

	conf := config.Get()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	api := router.Group("/api/v1")
	api.GET("/auth", ct.getUser)

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))

}
