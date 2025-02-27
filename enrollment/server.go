package main

import (
	"enrollment/repositories"
	"enrollment/services"

	"github.com/labstack/echo/v4"
)

func main() {
	repo := repositories.NewEnrollmentRepository()
	service := services.NewEnrollmentService(repo)

	e := echo.New()
	api := e.Group("/api")
	api.GET("/enrollments", service.GetEnrollments)
	api.GET("/enrollments/:id", service.GetEnrollment)
	api.POST("/enrollments", service.CreateEnrollment)

	e.Logger.Fatal(e.Start(":8083"))
}