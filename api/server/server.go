package server

import (
	"kraken/api-server/controllers"
	"net/http"

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
	cs.Options(sessions.Options{
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   3600 * 24 * 7,
		HttpOnly: false,
		Secure:   false,
	})
	r.Use(sessions.Sessions("kraken-session", cs))

	// Setup CORS policy
	r.Use(cors.New(cors.Config{
		// Add your testing environment here
		AllowOrigins: []string{
			"http://localhost",
			"http://localhost:8000",
			"http://100.85.100.49",
			"http://100.85.100.49:8000",
			"http://192.168.122.240",
			"http://192.168.122.240:8000",
			"http://192.168.64.3",
		},
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
	r.GET("/update/:todoid", controllers.Update)
	r.GET("/delete/:todoid", controllers.DeleteTodo)

	return Handler{router: r}
}

func (h *Handler) Run() error {
	err := h.router.Run("0.0.0.0:8888")
	return err
}
