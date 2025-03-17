package dto

type GetStudent struct {
	ID        int `json:"id"`
	Name 			string `json:"name"`
	Email     string `json:"email"`
}