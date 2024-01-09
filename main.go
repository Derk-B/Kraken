package main

import (
	api "KrakenSite/api"
	website "KrakenSite/website"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	api.Run(r)
	website.Run(r)

	r.Run()
}
