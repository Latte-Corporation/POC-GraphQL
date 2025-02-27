package main

import (
	"course/repositories"
	"course/services"

	"github.com/labstack/echo/v4"
)

func main() {

	courseRepository := repositories.NewCourseRepository()
	courseService := services.NewCourseService(courseRepository)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/courses", courseService.GetCourses)
	api.GET("/courses/:id", courseService.GetCourse)
	api.POST("/courses", courseService.SaveCourse)
	e.Logger.Fatal(e.Start(":8082"))
}
