package controllers

import (
	"context"
	"fmt"
	"kraken/api-server/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TodosData struct {
	Todos []models.Todo
}

type TodosResponse struct {
	Status string
	Data   TodosData
}

type TodoDetails struct {
	Title string
}

func GetAllTodos(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user")

	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	var todos []models.Todo

	err := models.DB.NewSelect().Model(&todos).Where("user_id = ?", userId).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	var todoData = new(TodosData)
	todoData.Todos = todos

	var todoResponse = new(TodosResponse)
	todoResponse.Data = *todoData
	todoResponse.Status = "Success"

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, todoResponse)
}

func AddTodo(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user")

	// UserId is always a positive number.
	// If it is -1 then no userid could be found an the user is unauthorized
	if userId == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Unauthorized",
			"message": "You are unauthorized",
		})
		return
	}

	// Get new todo from body
	todoDetails := new(TodoDetails)
	if err := c.BindJSON(&todoDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	todo := models.Todo{
		Title:     todoDetails.Title,
		Completed: false,
		UserId:    userId.(int64),
	}

	// Store new todo in DB
	_, err := models.DB.NewInsert().Model(&todo).Exec(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error while adding todo",
		})
		return
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo added successfully",
	})
}

func Update(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user")

	// Parse id of todo from the request
	reqTodoId, err := strconv.Atoi(c.Param("todoid"))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Error while updating todo",
		})
		return
	}

	// Delete todo from database
	todo := new(models.Todo)
	err = models.DB.NewSelect().Model(todo).Where("id = ?", reqTodoId).Where("user_id = ?", userId).Scan(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Error while updating todo",
		})
		return
	}

	todo.Completed = !todo.Completed
	_, err = models.DB.NewUpdate().Model(todo).Where("id = ?", reqTodoId).Where("user_id = ?", userId).Exec(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Error while updating todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo updated successfully",
	})
}

func DeleteTodo(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user")

	// Parse id of todo from the request
	reqTodoId, err := strconv.Atoi(c.Param("todoid"))

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Error while deleting todo",
		})
		return
	}

	// Delete todo from database
	todo := new(models.Todo)
	res, err := models.DB.NewDelete().Model(todo).Where("id = ?", reqTodoId).Where("user_id = ?", userId).Exec(context.Background())
	rowsAffected, err2 := res.RowsAffected()

	if err != nil || err2 != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Error while deleting todo",
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Unable to delete todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo deleted successfully",
	})
}
