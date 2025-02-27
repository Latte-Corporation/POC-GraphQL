package repositories

import (
	"enrollment/domains"
	"enrollment/dto"
)

type EnrollmentRepository interface {
	CreateEnrollment(enrollment *dto.PostEnrollment) (*domains.Enrollment, error)
	GetEnrollment(id int) (*domains.Enrollment, error)
	GetEnrollments() ([]*domains.Enrollment, error)
}

type enrollmentRepository struct {
	enrollments map[int]*domains.Enrollment
}

func NewEnrollmentRepository() EnrollmentRepository {
	return &enrollmentRepository{
		enrollments: make(map[int]*domains.Enrollment),
	}
}

func (r *enrollmentRepository) CreateEnrollment(enrollment *dto.PostEnrollment) (*domains.Enrollment, error) {
	id := len(r.enrollments) + 1
	newEnrollment := &domains.Enrollment{
		ID:        id,
		StudentID: enrollment.StudentID,
		CourseID:  enrollment.CourseID,
	}
	r.enrollments[id] = newEnrollment
	return newEnrollment, nil
}

func (r *enrollmentRepository) GetEnrollment(id int) (*domains.Enrollment, error) {
	return r.enrollments[id], nil
}

func (r *enrollmentRepository) GetEnrollments() ([]*domains.Enrollment, error) {
	var enrollments []*domains.Enrollment
	for _, enrollment := range r.enrollments {
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
