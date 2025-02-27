package main

import (
	"student/services"
	"student/repositories"
	echo "github.com/labstack/echo/v4"
)

func main() {
	studentRepository := repositories.NewStudentRepository()
	studentService := services.NewStudentService(studentRepository)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/students", studentService.GetStudents)
	api.GET("/students/:id", studentService.GetStudent)
	api.POST("/students", studentService.CreateStudent)
	e.Logger.Fatal(e.Start(":8081"))
}
