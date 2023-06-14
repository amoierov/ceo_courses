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

	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:63342"} // Разрешенные источники запросов
	//config.AllowOrigins = []string{"*"}                                       // Разрешенные источники запросов
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Разрешенные методы
	//config.AllowHeaders = []string{"Content-Type"}                            // Разрешенные заголовки
	//router.Use(cors.New(config))
	router.Use(cors.Default())
	//router.Use(corsHeader)
	router.GET("/user", h.index)
	router.GET("/admin", h.admin)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	router.GET("/courses/1/topics", h.getTopicsByCourse)

	api := router.Group("/api", h.userIdentity)
	{
		courses := api.Group("/courses")
		{
			courses.POST("/", h.createCourse)
			courses.GET("/", h.getAllCourses)
			courses.GET("/user", h.getCoursesByIdUser)
			courses.PUT("/:id", h.updateCourse)
			courses.DELETE("/:id", h.deleteCourse)
			courses.POST("/subscribe", h.subscribeUser)
			courses.GET("/author", h.getAuthors)

			topics := courses.Group(":id/topics")
			{
				topics.POST("/", h.createTopic)
				topics.GET("/", h.getAllTopics)
				topics.GET("/:topic_id", h.getTopicById)
				topics.PUT("/:topic_id", h.updateTopic)
				topics.DELETE("/:topic_id", h.deleteTopic)
			}
		}
		forms := api.Group("/forms")
		{
			forms.GET("/questions/:form_id", h.GetForm)
			forms.POST("/preferences", h.UserPreferences)
			forms.POST("/preferences/recommendation", h.createRecommendation)
		}

	}
	return router
}
