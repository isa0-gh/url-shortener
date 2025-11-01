package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isa0-gh/urlshorter/database"
)

func RedirectShortUrl(c *gin.Context) {
	id := c.Param("id")
	database.DB.Query("SELECT redirect_url FROM short_urls WHERE id = $1", id)

}
