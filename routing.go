package main

import (
	"github.com/Luthfiansyah/warpin-message/app/handlers"
	"github.com/gin-gonic/gin"
	"time"
)

var runMode string
var theTime time.Time

func initRoutes(debug bool) {

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// PING
	router.GET("/", Ping) // INIT SERVER RUN

	// API
	v1 := router.Group("/v1")
	v1.POST("/message", handlers.AddMessage)
	v1.GET("/message", handlers.GetAllMessage)
}
