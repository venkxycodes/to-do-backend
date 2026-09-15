package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do/auth"
	"to-do/config"
	"to-do/contract"
	"to-do/handler"
	"to-do/service"
)

// middleware are to be added

type Options struct {
	Conf         *config.Config
	Dependencies *service.ServerDependencies
}

func InitRouter(opts Options) *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware(opts.Conf.CORSOrigins))
	router.Use(requestSizeLimit(1 << 20))
	router.Use(gin.Recovery())
	contract.RegisterValidators()

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	userHandler := handler.NewUserHandler(opts.Dependencies.UserService)
	todoHandler := handler.NewToDoHandler(opts.Dependencies.ToDoService)
	InitUserRouter(router, &userHandler)
	InitToDoRouter(router, &todoHandler, opts.Conf.JWTSecret)
	return router
}

func InitToDoRouter(router *gin.Engine, handler *handler.ToDoHandler, jwtSecret string) {
	v1 := router.Group("to-do/v1", auth.RequireAuth(jwtSecret))
	v1.POST("task", handler.CreateTask)
	v1.PUT("task", handler.UpdateTask)
	v1.GET("tasks", handler.GetTasks)
	v1.PATCH("task", handler.UpdateTaskStatus)
}

func InitUserRouter(router *gin.Engine, handler *handler.UserHandler) {
	v1 := router.Group("to-do/v1/user")
	v1.POST("sign-up", handler.SignUpUser)
	v1.POST("login", handler.LoginUser)
}
