package main

import (
	"github.com/gin-gonic/gin"
	"kraken/web/web"
)

func main() {

	r := gin.Default()
	web.Run(r)

	err := r.Run()
	if err != nil {
		return
	}
}
