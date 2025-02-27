package services

import (
	"enrollment/dto"
	"enrollment/repositories"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type EnrollmentService interface {
	CreateEnrollment(c echo.Context) error
	GetEnrollment(c echo.Context) error
	GetEnrollments(c echo.Context) error
	GetEnrollmentsForStudent(c echo.Context) error
	GetEnrollmentsForCourse(c echo.Context) error
}

type enrollmentService struct {
	repo repositories.EnrollmentRepository
}

func NewEnrollmentService(repo repositories.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{
		repo: repo,
	}
}

func (s *enrollmentService) CreateEnrollment(c echo.Context) error {
	var enrollment dto.PostEnrollment
	if err := c.Bind(&enrollment); err != nil {
		c.JSON(400, err)
		return err
	}

	// Check if student exists
	studentResp, err := http.Get("http://localhost:8081/api/students/" + strconv.Itoa(enrollment.StudentID))
	if err != nil || studentResp.StatusCode != 200 {
		c.JSON(400, "Invalid student ID")
		return err
	}

	// Check if course exists
	courseResp, err := http.Get("http://localhost:8082/api/courses/" + strconv.Itoa(enrollment.CourseID))
	if err != nil || courseResp.StatusCode != 200 {
		c.JSON(400, "Invalid course ID")
		return err
	}

	newEnrollment, err := s.repo.CreateEnrollment(&enrollment)
	if err != nil {
		c.JSON(500, err)
		return err
	}

	return c.JSON(200, newEnrollment)
}

func (s *enrollmentService) GetEnrollment(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, "Invalid enrollment ID")
		return err
	}

	enrollment, err := s.repo.GetEnrollment(idInt)
	if err != nil {
		c.JSON(500, err)
		return err
	}

	return c.JSON(200, enrollment)
}

func (s *enrollmentService) GetEnrollments(c echo.Context) error {
	enrollments, err := s.repo.GetEnrollments()
	if err != nil {
		c.JSON(500, err)
	}
	if len(enrollments) == 0 {
		return c.JSON(404, "No enrollments found")
	}
	return c.JSON(200, enrollments)
}

func (s *enrollmentService) GetEnrollmentsForStudent(c echo.Context) error {
	studentID := c.Param("student_id")
	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		c.JSON(400, "Invalid student ID")
		return err
	}

	enrollments, err := s.repo.GetEnrollmentsForStudent(studentIDInt)
	if err != nil {
		c.JSON(500, err)
	}
	if len(enrollments) == 0 {
		return c.JSON(404, "No enrollments found")
	}
	return c.JSON(200, enrollments)
}

func (s *enrollmentService) GetEnrollmentsForCourse(c echo.Context) error {
	courseID := c.Param("course_id")
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		c.JSON(400, "Invalid course ID")
		return err
	}

	enrollments, err := s.repo.GetEnrollmentsForCourse(courseIDInt)
	if err != nil {
		c.JSON(500, err)
	}
	if len(enrollments) == 0 {
		return c.JSON(404, "No enrollments found")
	}
	return c.JSON(200, enrollments)
}
