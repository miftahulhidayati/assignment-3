package main

import (
	"assignment-3/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func backgroundTask() {
	for range time.Tick(15 * time.Second) {
		go controllers.UpdateJson()
	}
}

func main() {
	go backgroundTask()

	router := gin.Default()

	router.Static("/css", "./assets/css")
	router.Static("/js", "./assets/js")
	router.Static("/images", "./assets/images")
	router.Static("/font", "./assets/font")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", controllers.GetStatus)
	router.Run(":8080")
}
