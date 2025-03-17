package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/dto"
	"gateway/graph/model"
	"net/http"
	"os"
	"strconv"
)

type StudentRepository interface {
	GetStudent(id int) (*model.Student, error)
	GetStudents() ([]*model.Student, error)
	CreateStudent(input model.StudentInput) (*model.Student, error)
}

type studentRepository struct {
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{}
}

var studentURL = os.Getenv("STUDENT_SERVICE_URL")

func init() {
	if studentURL == "" {
		studentURL = "http://localhost:8081"
	}
}

func (r *studentRepository) GetStudent(id int) (*model.Student, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/students/%d", studentURL, id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch student: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var student dto.GetStudent
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, fmt.Errorf("failed to decode student: %w", err)
	}

	studentModel := model.Student{
		ID:    strconv.Itoa(student.ID),
		Name:  student.Name,
		Email: student.Email,
	}
	return &studentModel, nil
}

func (r *studentRepository) GetStudents() ([]*model.Student, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/students", studentURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch students: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var students []dto.GetStudent
	if err := json.NewDecoder(resp.Body).Decode(&students); err != nil {
		return nil, fmt.Errorf("failed to decode students: %w", err)
	}

	var studentModels []*model.Student
	for _, student := range students {
		studentModel := model.Student{
			ID:    strconv.Itoa(student.ID),
			Name:  student.Name,
			Email: student.Email,
		}
		studentModels = append(studentModels, &studentModel)
	}
	return studentModels, nil
}

func (r *studentRepository) CreateStudent(input model.StudentInput) (*model.Student, error) {
	studentData := dto.CreateStudent{
		Name:  input.Name,
		Email: input.Email,
	}

	jsonData, err := json.Marshal(studentData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal student data: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/students", studentURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create student: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var student dto.GetStudent
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, fmt.Errorf("failed to decode student: %w", err)
	}

	studentModel := model.Student{
		ID:    strconv.Itoa(student.ID),
		Name:  student.Name,
		Email: student.Email,
	}
	return &studentModel, nil
}
