package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func HandleRequests(port string)  {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.NoRoute(controllers.RouterNotFound)

	r.GET("/salutation/:name", controllers.Salutation)

	r.GET("/index", controllers.ShowIndexPage)

	r.GET("/students", controllers.ShowStudents)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCpf)
	r.POST("/students", controllers.InsertStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.Run(port)
}
