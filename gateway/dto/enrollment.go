package dto

type GetEnrollment struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}

type PostEnrollment struct {
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}