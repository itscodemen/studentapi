package router

import (
	"proj1/api/handlers"

	"flag"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/student", handlers.CreateStudent)
	r.GET("/student", handlers.GetStudents)
	r.GET("/student/:id", handlers.GetStudentByID)
	r.PUT("/student/:id", handlers.UpdateStudent)
	r.DELETE("/student/:id", handlers.DeleteStudent)

	prt := flag.String("port", ":8080", "Port Number")
	flag.Parse()
	port := *prt
	r.Run(":" + port)
	return r
}
