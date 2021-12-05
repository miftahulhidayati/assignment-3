package controllers

import (
	"assignment-3/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	var status models.Status
	var statusDescription models.StatusDescription
	// var statusShow models.StatusShow
	min := 1
	max := 20
	// for range time.Tick(15 * time.Second) {
		status.Water = rand.Intn(max - min)
		status.Wind = rand.Intn(max - min)
	// }
	
	if status.Water <= 5 {
		statusDescription.Water = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		statusDescription.Water = "Siaga"
	} else if status.Water > 8 {
		statusDescription.Water = "Bahaya"
	}

	if status.Wind <= 5 {
		statusDescription.Wind = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		statusDescription.Wind = "Siaga"
	} else if status.Wind > 15 {
		statusDescription.Wind = "Bahaya"
	}
	// statusShow = models.StatusShow{
	// 	Status:            status,
	// 	StatusDescription: statusDescription,
	// }

	// data, _ := json.Marshal(statusDescription)

	c.HTML(http.StatusOK, "template.html", statusDescription)
}