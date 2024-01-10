package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(engine *gin.Engine) {
	r := engine.Group("/api")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
