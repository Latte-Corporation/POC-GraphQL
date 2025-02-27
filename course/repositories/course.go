package repositories

import (
	"course/domains"
	"course/dto"
)

type CourseRepository interface {
	CreateCourse(course *dto.PostCourse) (*domains.Course, error)
	GetCourse(id int) (*domains.Course, error)
	GetCourses() ([]*domains.Course, error)
}

type courseRepository struct {
	courses map[int]*domains.Course
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{
		courses: make(map[int]*domains.Course),
	}
}

func (r *courseRepository) CreateCourse(course *dto.PostCourse) (*domains.Course, error) {
	id := len(r.courses) + 1
	newCourse := &domains.Course{
		ID:          id,
		Title:       course.Title,
		Description: course.Description,
	}
	r.courses[id] = newCourse
	return newCourse, nil
}

func (r *courseRepository) GetCourse(id int) (*domains.Course, error) {
	return r.courses[id], nil
}

func (r *courseRepository) GetCourses() ([]*domains.Course, error) {
	var courses []*domains.Course
	for _, course := range r.courses {
		courses = append(courses, course)
	}
	return courses, nil
}
