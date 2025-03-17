package repositories

import (
	"encoding/json"
	"fmt"
	"gateway/dto"
	"gateway/graph/model"
	"net/http"
	"os"
	"strconv"
)

type CourseRepository interface {
	GetCourse(id int) (*model.Course, error)
	GetCourses() ([]*model.Course, error)
}

type courseRepository struct {
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{}
}

var courseURL = os.Getenv("COURSE_SERVICE_URL")

func init() {
	if courseURL == "" {
		courseURL = "http://localhost:8082"
	}
}

func (r *courseRepository) GetCourse(id int) (*model.Course, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/courses/%d", courseURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch course: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var course dto.GetCourse
	if err := json.NewDecoder(resp.Body).Decode(&course); err != nil {
		return nil, fmt.Errorf("failed to decode course: %w", err)
	}

	courseModel := model.Course{
		ID:          strconv.Itoa(course.ID),
		Title:       course.Title,
		Description: &course.Description,
	}
	return &courseModel, nil
}

func (r *courseRepository) GetCourses() ([]*model.Course, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/courses", courseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch courses: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var courses []dto.GetCourse
	if err := json.NewDecoder(resp.Body).Decode(&courses); err != nil {
		return nil, fmt.Errorf("failed to decode courses: %w", err)
	}

	var courseModels []*model.Course
	for _, course := range courses {
		courseModel := model.Course{
			ID:          strconv.Itoa(course.ID),
			Title:       course.Title,
			Description: &course.Description,
		}
		courseModels = append(courseModels, &courseModel)
	}
	return courseModels, nil
}
