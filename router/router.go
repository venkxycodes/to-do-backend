package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do/config"
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
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	todoHandler := handler.NewToDoHandler(opts.Dependencies.ToDoService)
	InitTodoRouter(router, &todoHandler)

	return router
}

func InitTodoRouter(router *gin.Engine, handler *handler.ToDoHandler) {
	v1 := router.Group("api/v1")
	v1.POST("reminder", handler.CreateToDo)
}
