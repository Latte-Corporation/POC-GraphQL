package repositories

import (
    "encoding/json"
    "fmt"
    "gateway/dto"
		"gateway/domains"
    "net/http"
    "os"
)

type EnrollmentRepository interface {
    GetEnrollment(id int) (*domains.Enrollment, error)
    GetEnrollments() ([]*domains.Enrollment, error)
    GetEnrollmentsByStudentID(studentID int) ([]*domains.Enrollment, error)
    GetEnrollmentsByCourseID(courseID int) ([]*domains.Enrollment, error)
}

type enrollmentRepository struct {
}

func NewEnrollmentRepository() EnrollmentRepository {
    return &enrollmentRepository{}
}

var enrollmentURL = os.Getenv("ENROLLMENT_SERVICE_URL")

const errorMessage = "failed to fetch enrollments"

func (r *enrollmentRepository) GetEnrollment(id int) (*domains.Enrollment, error) {
    resp, err := http.Get(fmt.Sprintf("%s/api/enrollments/%d", enrollmentURL, id))
    if err != nil {
        return nil, fmt.Errorf("%s: %w",errorMessage, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("%s: %d",errorMessage, resp.StatusCode)
    }

    var enrollment dto.GetEnrollment
    if err := json.NewDecoder(resp.Body).Decode(&enrollment); err != nil {
        return nil, fmt.Errorf("failed to decode enrollment: %w", err)
    }

    enrollmentModel := domains.Enrollment{
        ID:        enrollment.ID,
        StudentID: enrollment.StudentID,
        CourseID:  enrollment.CourseID,
    }
    return &enrollmentModel, nil
}

func (r *enrollmentRepository) GetEnrollments() ([]*domains.Enrollment, error) {
    resp, err := http.Get(fmt.Sprintf("%s/api/enrollments", enrollmentURL))
    if err != nil {
				return nil, fmt.Errorf("%s: %w", errorMessage, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var enrollments []dto.GetEnrollment
    if err := json.NewDecoder(resp.Body).Decode(&enrollments); err != nil {
        return nil, fmt.Errorf("%s: %w", errorMessage, err)
    }

    var enrollmentModels []*domains.Enrollment
    for _, enrollment := range enrollments {
        enrollmentModel := domains.Enrollment{
            ID:        enrollment.ID,
            StudentID: enrollment.StudentID,
            CourseID:  enrollment.CourseID,
        }
        enrollmentModels = append(enrollmentModels, &enrollmentModel)
    }
    return enrollmentModels, nil
}

func (r *enrollmentRepository) GetEnrollmentsByStudentID(studentID int) ([]*domains.Enrollment, error) {
    resp, err := http.Get(fmt.Sprintf("%s/api/enrollments/students/%d", enrollmentURL, studentID))
    if err != nil {
        return nil, fmt.Errorf("failed to fetch enrollments: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var enrollments []dto.GetEnrollment
    if err := json.NewDecoder(resp.Body).Decode(&enrollments); err != nil {
        return nil, fmt.Errorf("failed to decode enrollments: %w", err)
    }

    var enrollmentModels []*domains.Enrollment
    for _, enrollment := range enrollments {
        enrollmentModel := domains.Enrollment{
            ID:        enrollment.ID,
            StudentID: enrollment.StudentID,
            CourseID:  enrollment.CourseID,
        }
        enrollmentModels = append(enrollmentModels, &enrollmentModel)
    }
    return enrollmentModels, nil
}

func (r *enrollmentRepository) GetEnrollmentsByCourseID(courseID int) ([]*domains.Enrollment, error) {
    resp, err := http.Get(fmt.Sprintf("%s/api/enrollments/courses/%d", enrollmentURL, courseID))
    if err != nil {
        return nil, fmt.Errorf("failed to fetch enrollments: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    var enrollments []dto.GetEnrollment
    if err := json.NewDecoder(resp.Body).Decode(&enrollments); err != nil {
        return nil, fmt.Errorf("failed to decode enrollments: %w", err)
    }

    var enrollmentModels []*domains.Enrollment
    for _, enrollment := range enrollments {
        enrollmentModel := domains.Enrollment{
            ID:        enrollment.ID,
            StudentID: enrollment.StudentID,
            CourseID:  enrollment.CourseID,
        }
        enrollmentModels = append(enrollmentModels, &enrollmentModel)
    }
    return enrollmentModels, nil
}