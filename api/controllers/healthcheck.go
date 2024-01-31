package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var ApiVersion string
var hostname = "NO_HOSTNAME"

func Ping(c *gin.Context) {
	hostname, _ = os.Hostname()
	if ApiVersion == "" {
		ApiVersion = "DEBUG_BUILD"
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Kraken API is running on host: " + hostname,
		"version": ApiVersion,
	})
}
