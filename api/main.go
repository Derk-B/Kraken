package main

import (
	"github.com/gin-gonic/gin"
	api "kraken/api-server/endpoint"
)

func main() {

	r := gin.Default()

	api.Run(r)

	err := r.Run()
	if err != nil {
		return
	}
}
