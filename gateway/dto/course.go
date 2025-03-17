package dto

type GetCourse struct {
	ID          int    `json:"id"`
	Title        string `json:"title"`
	Description string `json:"description"`
}

type CreateCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
