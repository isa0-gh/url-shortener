package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isa0-gh/url-shortener/database/utils"
	"github.com/isa0-gh/url-shortener/models"
)

func RedirectShortUrl(c *gin.Context) {
	id := c.Param("id")
	url, _ := utils.GetUrl(id)
	if url == "" {
		c.Redirect(302, "/") // redirect to index page
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
	resp, err := utils.Create(json.Url, json.Expire)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func DeleteShortUrl(c *gin.Context) {
	id := c.Param("id")
	utils.Delete(id)
	c.Redirect(302, "/")
}
