package main

import (
	"fmt"
	"os"

	"github.com/isa0-gh/urlshorter/database"

	"github.com/gin-gonic/gin"
	"github.com/isa0-gh/urlshorter/routes"
)

func main() {
	var port string = os.Getenv("PORT")
	var bindAddress string
	if !port {
		fmt.Println("Set a port to run! example: PORT=8080 ./urlshorter")
		return
	} else {
		bindAddress = fmt.Sprintf(":%s", port)
	}
	database.Init()

	router := gin.Default()
	router.GET("/s/:id", routes.RedirectShortUrl)
	router.POST("/api/create", routes.CreateNewShortUrl)
	router.Run(bindAddress)
}
