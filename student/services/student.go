package services

import (
	"strconv"
	"student/dto"
	"student/repositories"

	echo "github.com/labstack/echo/v4"
)

type StudentService interface {
	CreateStudent(c echo.Context) error
	GetStudent(c echo.Context) error
	GetStudents(c echo.Context) error
}

type studentService struct {
	repo repositories.StudentRepository
}

func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) CreateStudent(c echo.Context) error{
	var student dto.PostStudent
	if err := c.Bind(&student); err != nil {
		c.JSON(400, err)
		return err
	}

	newStudent, err := s.repo.CreateStudent(&student) 
	if err != nil {
		c.JSON(500, err)
		return err
	}

	return c.JSON(200, newStudent)
}

func (s *studentService) GetStudent(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, "Invalid student ID")
		return err
	}

	student, err := s.repo.GetStudent(idInt)
	if err != nil {
		c.JSON(500, err)
		return err
	}

	return c.JSON(200, student)
}

func (s *studentService) GetStudents(c echo.Context) error {
	students, err := s.repo.GetStudents()
	if err != nil {
		c.JSON(500, err)
		return err
	}
	if len(students) == 0 {
		c.JSON(404, "No students found")
		return nil
	}
	return c.JSON(200, students)
}

