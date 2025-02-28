package main

import (
	"course/repositories"
	"course/services"
	"os"

	"github.com/labstack/echo/v4"
)

const defaultPort = "8082"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	courseRepository := repositories.NewCourseRepository()
	courseService := services.NewCourseService(courseRepository)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/courses", courseService.GetCourses)
	api.GET("/courses/:id", courseService.GetCourse)
	api.POST("/courses", courseService.SaveCourse)
	e.Logger.Fatal(e.Start(":"+port))
}
