package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isa0-gh/urlshorter/database/utils"
	"github.com/isa0-gh/urlshorter/models"
)

func RedirectShortUrl(c *gin.Context) {
	id := c.Param("id")
	url, err := utils.GetUrl(id)
	if err != nil {
		c.Redirect(404, "/") // redirect to index page if its invalid url id
		return
	}

	c.Redirect(302, url) // redirect to original url

}

func CreateNewShortUrl(c *gin.Context) {
	var json models.Body
	var resp models.NewUrl
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Body"})
		return
	}
	resp, err := utils.Create(json.lUrl, json.Expire)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
