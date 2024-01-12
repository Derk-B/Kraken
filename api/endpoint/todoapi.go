package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:blogposts`

	Id          int64 `bun:",pk,autoincrement"`
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

func Run(r *gin.Engine, db *bun.DB) {

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
		userid := 0 // TODO: Retrieve actual userid

		// Userid is always a positive number.
		// If it is -1 then then no userid could be found an the user is unauthorized
		if userid == -1 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "Unauthorized",
				"message": "You are unauthorized",
			})
			return
		}

		// Get new todo from body
		title := c.PostForm("title")
		desc := c.PostForm("description")

		todo := Todo{
			Title:       title,
			Description: desc,
			Userid:      userid,
		}

		// Store new todo in DB
		_, err := db.NewInsert().Model(&todo).Exec(context.Background())

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"status":  "Fail",
				"message": "Error while adding todo",
			})
			return
		}

		// Send response
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Todo added successfully",
		})
	})

	r.DELETE("/todo/:todoid", func(c *gin.Context) {
		userid := 0 // TODO: Retrieve actual userid

		// Parse id of todo from the request
		req_todoid, err := strconv.Atoi(c.Param("todoid"))

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"status":  "Fail",
				"message": "Error while deleting todo",
			})
			return
		}

		// Delete todo from database
		todo := new(Todo)
		res, err := db.NewDelete().Model(todo).Where("id = ?", req_todoid).Where("userId = ?", userid).Exec(context.Background())

		rowsAffected, err2 := res.RowsAffected()

		if err != nil || err2 != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"status":  "Fail",
				"message": "Error while deleting todo",
			})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "Fail",
				"message": "Unable to delete todo",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Todo deleted successfully",
		})
	})
}
