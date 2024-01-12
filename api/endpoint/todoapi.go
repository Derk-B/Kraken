package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Userid      int
}

type TodosData struct {
	Todos []Todo
}

type TodosResponse struct {
	Status string
	Data   TodosData
}

func Run(r *gin.Engine) {
	// Get env variables for connection string
	envFile, err := godotenv.Read("../.env")

	if err != nil {
		fmt.Println("Could not load env file")
	}

	PSQL_PORT := envFile["PSQL_PORT"]
	PSQL_USERNAME := envFile["PSQL_USERNAME"]
	PSQL_PASSWORD := envFile["PSQL_PASSWORD"]
	PSQL_DATABASE := envFile["PSQL_DATABASE"]

	dsn := "postgres://" + PSQL_USERNAME + ":" + PSQL_PASSWORD + "@localhost:" + PSQL_PORT + "/" + PSQL_DATABASE + "?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	defer sqldb.Close()
	defer context.Background().Done()

	db := bun.NewDB(sqldb, pgdialect.New())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todos", func(c *gin.Context) {
		userId := c.Query("userid")

		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		var todos []Todo

		err := db.NewSelect().Model(&todos).Where("userid = ?", userId).Scan(context.Background())

		if err != nil {
			fmt.Println(err)
		}

		var todo_data = new(TodosData)
		todo_data.Todos = todos

		var todo_response = new(TodosResponse)
		todo_response.Data = *todo_data
		todo_response.Status = "Success"

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, todo_response)
	})

	r.POST("/todo", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.DELETE("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
