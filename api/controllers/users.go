package controllers

import (
	"context"
	"fmt"
	"kraken/api-server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SignUpDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailDetails struct {
	Email string `json:"email"`
}

func SignUp(c *gin.Context) {
	var signUpDetails SignUpDetails
	if err := c.BindJSON(&signUpDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: signUpDetails.Username,
		Email:    signUpDetails.Email,
		Password: signUpDetails.Password,
	}

	// Checking if already signed up with the email
	registeredUser := new(models.User)
	err := models.DB.NewSelect().Model(registeredUser).Where("email = ?", signUpDetails.Email).Scan(context.Background())
	if err != nil {
		registeredUser = nil
	}
	if registeredUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Account already exists.",
		})
		return
	}

	// Actual registration
	_, err = models.DB.NewInsert().Model(&user).Exec(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Internal server error.",
		})
		return
	}
	// Send response
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Account successfully created.",
	})
}

func SignIn(c *gin.Context) {
	fmt.Println("Hall ologin")
	var signInDetails SignInDetails
	if err := c.BindJSON(&signInDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	registeredUser := new(models.User)
	err := models.DB.NewSelect().Model(registeredUser).Where("email = ?", signInDetails.Email).Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if signInDetails.Password == registeredUser.Password {
		session := sessions.Default(c)
		session.Set("user", registeredUser.Username)
		session.Options(sessions.Options{
			MaxAge: 3600,
		})
		res := session.Get("user")
		fmt.Println(res)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to save session.",
			})
			return
		} else {
			session := sessions.Default(c)
			fmt.Println(session)
			res := session.Get("user")
			fmt.Println(res)
			fmt.Print("This is the user")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Password is incorrect.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Successfully authenticated user",
	})
}

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	status := session.Get("user")

	if status == nil || status == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid session token",
		})
		return
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to save session",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func ResetPassword(c *gin.Context) {
	var emailDetails EmailDetails
	if err := c.BindJSON(&emailDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Implement password reset logic...
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	status := session.Get("user")
	fmt.Println(session)
	fmt.Println(status)
	if status == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "unauthorized",
		})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
