package services

import (
	"fmt"
	"gateway/graph/model"
	"gateway/repositories"
)

type EnrollmentService interface {
	GetCoursesByStudentID(studentID int) ([]*model.Course, error)
	GetStudentsByCourseID(courseID int) ([]*model.Student, error)
}

type enrollmentService struct {
	enrollmentRepo repositories.EnrollmentRepository
	studentRepo    repositories.StudentRepository
	courseRepo     repositories.CourseRepository
}

func NewEnrollmentService(enrollmentRepo repositories.EnrollmentRepository, studentRepo repositories.StudentRepository, courseRepo repositories.CourseRepository) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		studentRepo:    studentRepo,
		courseRepo:     courseRepo,
	}
}

func (s *enrollmentService) GetCoursesByStudentID(studentID int) ([]*model.Course, error) {
	enrollments, err := s.enrollmentRepo.GetEnrollmentsByStudentID(studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch courses by student ID: %w", err)
	}

	var courses []*model.Course
	for _, enrollment := range enrollments {
		course, err := s.courseRepo.GetCourse(enrollment.CourseID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch course by ID: %w", err)
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (s *enrollmentService) GetStudentsByCourseID(courseID int) ([]*model.Student, error) {
	enrollments, err := s.enrollmentRepo.GetEnrollmentsByCourseID(courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch students by course ID: %w", err)
	}

	var students []*model.Student
	for _, enrollment := range enrollments {
		student, err := s.studentRepo.GetStudent(enrollment.StudentID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch student by ID: %w", err)
		}

		students = append(students, student)
	}

	return students, nil
}