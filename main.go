package main

// Importing fmt and time
import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

// type Status struct {
// 	Water int `json:"water"`
// 	Wind  int `json:"wind"`
// }
// type StatusDescription struct{
// 	Water string `json:"water"`
// 	Wind string `json:"wind"`
// }
// type StatusShow struct{
// 	Status Status `json:"status"`
// 	StatusDescription StatusDescription `json:"status_description"`
// }

// Golang program to illustrate the usage of
// Tick() function

// Main function
func main() {

	// Calling Tick method
	// using range keyword

		// min := 1
		// max := 20
		// status := Status{
		// 	Water: rand.Intn(max-min), 
		// 	Wind: rand.Intn(max-min),
		// }
		// var statusDescription StatusDescription
		// if status.Water <= 5 {
		// 	statusDescription.Water = "Aman"
		// } else if ( status.Water >= 6 && status.Water <= 8){
		// 	statusDescription.Water = "Siaga"
		// } else if (status.Water > 8){
		// 	statusDescription.Water = "Bahaya"
		// }

		// if status.Wind <= 5 {
		// 	statusDescription.Wind = "Aman"
		// } else if ( status.Wind >= 7 && status.Wind <= 15){
		// 	statusDescription.Wind = "Siaga"
		// } else if (status.Wind > 15){
		// 	statusDescription.Wind = "Bahaya"
		// }
		// statusShow := StatusShow{
		// 	Status: status,
		// 	StatusDescription: statusDescription,
		// }

		// data, _ := json.Marshal(statusShow)

		router := gin.Default()

		router.Static("/css", "./assets/css")
		router.Static("/js", "./assets/js")
		router.Static("/images", "./assets/images")
		router.Static("/font", "./assets/font")
		router.LoadHTMLGlob("templates/*.html")
		router.GET("/index", controllers.GetStatus)
		router.Run(":8080")
	
}
