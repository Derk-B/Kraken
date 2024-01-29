package endpoint

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"kraken/api-server/controllers"
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
		AllowOrigins: []string{"http://192.168.122.240:8000"},
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

	// Setup todos routes
	r.Use(controllers.AuthRequired).GET("/todos", controllers.GetAllTodos)
	r.Use(controllers.AuthRequired).POST("/todo", controllers.AddTodo)
	r.Use(controllers.AuthRequired).DELETE("/todo/:todoid", controllers.DeleteTodo)

	r.GET("/ping", controllers.Ping)

	return Handler{router: r}
}

func (h *Handler) Run() {
	err := h.router.Run("0.0.0.0:8888")
	if err != nil {
		return
	}
}
