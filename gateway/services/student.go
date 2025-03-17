package services

import (
	"fmt"
	"gateway/graph/model"
	"gateway/repositories"
	"strconv"
)

type StudentService interface {
	GetStudent(id int) (*model.Student, error)
	GetStudents() ([]*model.Student, error)
	CreateStudent(input model.StudentInput) (*model.Student, error)
}

type studentService struct {
	repo repositories.StudentRepository
	enrollmentService EnrollmentService
}

func NewStudentService(repo repositories.StudentRepository, enrollmentService EnrollmentService) StudentService {
	return &studentService{
		repo: repo,
		enrollmentService: enrollmentService,
	}
}

func (s *studentService) GetStudent(id int) (*model.Student, error) {
	student, err := s.repo.GetStudent(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch student: %w", err)
	}

	courses, err := s.enrollmentService.GetCoursesByStudentID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch courses for student: %w", err)
	}

	student.Courses = courses
	return student, nil
}

func (s *studentService) GetStudents() ([]*model.Student, error) {
	students, err := s.repo.GetStudents()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch students: %w", err)
	}

	for _, student := range students {
		studentID, err := strconv.Atoi(student.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid student ID: %w", err)
		}
		courses, err := s.enrollmentService.GetCoursesByStudentID(studentID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch courses for student: %w", err)
		}

		student.Courses = courses
	}

	return students, nil
}

func (s *studentService) CreateStudent(input model.StudentInput) (*model.Student, error) {
	student, err := s.repo.CreateStudent(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create student: %w", err)
	}
	return student, nil
}
