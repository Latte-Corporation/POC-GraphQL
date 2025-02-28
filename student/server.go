package main

import (
	"os"
	"student/repositories"
	"student/services"

	echo "github.com/labstack/echo/v4"
)

const defaultPort = "8081"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	studentRepository := repositories.NewStudentRepository()
	studentService := services.NewStudentService(studentRepository)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/students", studentService.GetStudents)
	api.GET("/students/:id", studentService.GetStudent)
	api.POST("/students", studentService.CreateStudent)
	e.Logger.Fatal(e.Start(":"+port))
}
