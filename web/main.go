package main

import (
	"kraken/web/web"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	web.Run(r)

	err := r.Run()
	if err != nil {
		return
	}
}
