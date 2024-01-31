package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "NO_HOSTNAME"
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Kraken API is running on host: " + hostname,
	})
}
