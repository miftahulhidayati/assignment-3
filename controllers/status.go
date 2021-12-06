package controllers

import (
	"assignment-3/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {

	file, _ := ioutil.ReadFile("status.json")
	data := models.StatusShow{}
	_ = json.Unmarshal([]byte(file), &data)

	c.HTML(http.StatusOK, "template.html", data)
}
func UpdateJson() {
	var status models.Status
	var statusDescription models.StatusDescription
	var statusShow models.StatusShow
	min := 1
	max := 10
	status.Water = rand.Intn(max - min)
	status.Wind = rand.Intn(max - min)

	if status.Water <= 5 {
		statusDescription.Water = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		statusDescription.Water = "Siaga"
	} else if status.Water > 8 {
		statusDescription.Water = "Bahaya"
	}

	if status.Wind <= 6 {
		statusDescription.Wind = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		statusDescription.Wind = "Siaga"
	} else if status.Wind > 15 {
		statusDescription.Wind = "Bahaya"
	}
	statusShow = models.StatusShow{
		Status:            status,
		StatusDescription: statusDescription,
	}
	bytesData, _ := json.Marshal(statusShow)
	_ = ioutil.WriteFile("status.json", bytesData, 0644)
}
