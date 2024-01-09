package KrakenSite

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(engine *gin.Engine) {
	r := engine.Group("/api")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
