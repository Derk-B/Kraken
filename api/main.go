package main

import (
	api "kraken/api-server/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	api.Run(r)

	err := r.Run("localhost:8888")
	if err != nil {
		return
	}
}
