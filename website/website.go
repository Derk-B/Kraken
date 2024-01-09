package KrakenSite

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string
	Author string
	Price  float32
}

func Run(engine *gin.Engine) {

	r := engine.Group("/")

	books := []Book{
		{
			Title:  "Title 1",
			Author: "Me",
			Price:  49.99,
		},
		{
			Title:  "Title 2",
			Author: "Me",
			Price:  69.99,
		},
	}

	engine.LoadHTMLGlob("website/templates/*")

	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"books": books,
		})
	})
}
