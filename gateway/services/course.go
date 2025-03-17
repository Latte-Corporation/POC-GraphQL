package services

import (
	"fmt"
	"gateway/graph/model"
	"gateway/repositories"
	"strconv"
)

type CourseService interface {
	GetCourse(id int) (*model.Course, error)
	GetCourses() ([]*model.Course, error)
	CreateCourse(input model.CourseInput) (*model.Course, error)
}

type courseService struct {
	repo repositories.CourseRepository
	enrollmentService EnrollmentService
}

func NewCourseService(repo repositories.CourseRepository, enrollenrollmentService EnrollmentService) CourseService {
	return &courseService{
		repo: repo,
		enrollmentService: enrollenrollmentService,
	}
}

func (s *courseService) GetCourse(id int) (*model.Course, error) {
	course, err := s.repo.GetCourse(id)
	if err != nil {
			return nil, fmt.Errorf("failed to fetch course: %w", err)
	}

	students, err := s.enrollmentService.GetStudentsByCourseID(id)
	if err != nil {
			return nil, fmt.Errorf("failed to fetch students for course: %w", err)
	}

	course.Students = students
	return course, nil
}

func (s *courseService) GetCourses() ([]*model.Course, error) {
	courses, err := s.repo.GetCourses()
	if err != nil {
			return nil, fmt.Errorf("failed to fetch courses: %w", err)
	}

	for _, course := range courses {
		courseID, err := strconv.Atoi(course.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid course ID: %w", err)
		}
		students, err := s.enrollmentService.GetStudentsByCourseID(courseID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch students for course: %w", err)
		}

		course.Students = students
	}

	return courses, nil
}

func (s *courseService) CreateCourse(input model.CourseInput) (*model.Course, error) {
	course, err := s.repo.CreateCourse(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create course: %w", err)
	}
	return course, nil
}
