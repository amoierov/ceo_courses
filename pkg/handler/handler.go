package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"seo_courses/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}        // Разрешенные источники запросов
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"} // Разрешенные методы
	config.AllowHeaders = []string{"Content-Type"}                 // Разрешенные заголовки
	router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		courses := api.Group("/courses")
		{
			courses.POST("/", h.createCourse)
			courses.GET("/", h.getAllCourses)
			courses.GET("/:id", h.getCourseById)
			courses.PUT("/:id", h.updateCourse)
			courses.DELETE("/:id", h.deleteCourse)

			topics := courses.Group(":id/topics")
			{
				topics.POST("/", h.createTopic)
				topics.GET("/", h.getAllTopics)
				topics.GET("/:topic_id", h.getTopicById)
				topics.PUT("/:topic_id", h.updateTopic)
				topics.DELETE("/:topic_id", h.deleteTopic)
			}
		}
	}
	return router
}
