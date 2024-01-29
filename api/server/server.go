package server

import (
	"kraken/api-server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Handler struct {
	router *gin.Engine
	db     *bun.DB
}

func New() Handler {
	r := gin.Default()

	// Setup the cookie store for session management
	var secret = []byte("secret")
	cs := cookie.NewStore(secret)
	cs.Options(sessions.Options{Path: ""})
	r.Use(sessions.Sessions("kraken-session", cs))

	// Setup CORS policy
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost", "http://100.85.100.49:8000"},
		//AllowOrigins:        []string{"*"},
		AllowPrivateNetwork: true,
		AllowCredentials:    true,
		AllowWildcard:       true,
	}))

	// Setup auth routes
	r.POST("/user/signup", controllers.SignUp)
	r.POST("/user/signin", controllers.SignIn)
	r.GET("/user/signout", controllers.SignOut)
	r.POST("/user/password-reset", controllers.ResetPassword)

	r.GET("/ping", controllers.Ping)

	// Setup todos routes
	r.Use(controllers.AuthRequired)
	r.GET("/todos", controllers.GetAllTodos)
	r.POST("/todo", controllers.AddTodo)
	r.DELETE("/todo/:todoid", controllers.DeleteTodo)

	return Handler{router: r}
}

func (h *Handler) Run() {
	err := h.router.Run("0.0.0.0:8888")
	if err != nil {
		return
	}
}
