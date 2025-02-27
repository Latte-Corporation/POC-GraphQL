package repositories

import (
	"student/domains"
	"student/dto"
)
type StudentRepository interface {
	CreateStudent(student *dto.PostStudent) (*domains.Student, error)
	GetStudent(id int) (*domains.Student, error)
	GetStudents() ([]*domains.Student, error)
}

type studentRepository struct {
	students map[int]*domains.Student
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{
		students: make(map[int]*domains.Student),
	}
}

func (r *studentRepository) CreateStudent(student *dto.PostStudent) (*domains.Student, error) {
	id := len(r.students) + 1
	newStudent := &domains.Student{
		ID:        id,
		Name:      student.Name,
		Email:     student.Email,
	}
	r.students[1] = newStudent
	return newStudent, nil
}

func (r *studentRepository) GetStudent(id int) (*domains.Student, error) {
	return r.students[id], nil
}

func (r *studentRepository) GetStudents() ([]*domains.Student, error) {
	var students []*domains.Student
	for _, student := range r.students {
		students = append(students, student)
	}
	return students, nil
}

