package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"todolist/internal/service"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTaskByID)
			tasks.POST("/", h.createTask)
			tasks.DELETE("/:id", h.deleteTask)
			tasks.PUT("/:id", h.updateTask)
		}
	}
	return router
}
