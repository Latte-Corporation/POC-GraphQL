package services

import (
	"course/dto"
	"course/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CourseService interface {
	GetCourse(c echo.Context) error
	GetCourses(c echo.Context) error
	SaveCourse(c echo.Context) error
}

type courseService struct {
	repo repositories.CourseRepository
}

func NewCourseService(repo repositories.CourseRepository) CourseService {
	return &courseService{
		repo: repo,
	}
}

func (s *courseService) GetCourse(c echo.Context) error {
	id := c.Param("id")
	println("Getting course with ID: ", id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, "Invalid course ID")
		return err
	}

	course, err := s.repo.GetCourse(idInt)
	if err != nil {
		c.JSON(500, err)
		return err
	}

	c.JSON(200, course)
	return nil
}

func (s *courseService) GetCourses(c echo.Context) error {
	courses, err := s.repo.GetCourses()
	if err != nil {
		c.JSON(500, err)
		return err
	}

	c.JSON(200, courses)
	return nil
}

func (s *courseService) SaveCourse(c echo.Context) error {
	var course dto.PostCourse
	if err := c.Bind(&course); err != nil {
		c.JSON(400, err)
		return err
	}

	newCourse, err := s.repo.CreateCourse(&course)
	if err != nil {
		c.JSON(500, err)
		return err
	}

	println("New course created: ", newCourse)
	c.JSON(201, newCourse)
	return nil
}
