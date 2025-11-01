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
	if port == "" {
		fmt.Println("Default PORT(8080) will be used. To change port: PORT=500 ./urlshorter")
		bindAddress = fmt.Sprintf(":%s", "8080")
	} else {
		bindAddress = fmt.Sprintf(":%s", port)
	}
	database.Init()

	router := gin.Default()
	router.StaticFile("/", "./static/index.html")
	router.GET("/s/:id", routes.RedirectShortUrl)
	router.POST("/api/create", routes.CreateNewShortUrl)
	router.GET("/d/:id", routes.DeleteShortUrl)
	
	router.Run(bindAddress)
}
