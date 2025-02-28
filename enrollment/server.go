package main

import (
	"enrollment/repositories"
	"enrollment/services"
	"os"

	"github.com/labstack/echo/v4"
)

const defaultPort = "8083"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo := repositories.NewEnrollmentRepository()
	service := services.NewEnrollmentService(repo)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/enrollments", service.GetEnrollments)
	api.GET("/enrollments/:id", service.GetEnrollment)
	api.GET("/enrollments/students/:student_id", service.GetEnrollmentsForStudent)
	api.GET("/enrollments/courses/:course_id", service.GetEnrollmentsForCourse)
	api.POST("/enrollments", service.CreateEnrollment)

	e.Logger.Fatal(e.Start(":"+port))
}
