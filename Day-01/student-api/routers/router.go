package routers

import (
	"github.com/gin-gonic/gin"
	"student-api/handlers"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/healthcheck", handlers.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/students", handlers.CreateStudent)
		v1.GET("/students", handlers.GetStudents)
		v1.GET("/students/:id", handlers.GetStudent)
		v1.PUT("/students/:id", handlers.UpdateStudent)
		v1.DELETE("/students/:id", handlers.DeleteStudent)
	}

	return router
}