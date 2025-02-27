package dto

type PostStudent struct {
	Name 			string `json:"name"`
	Email     string `json:"email"`
}

type GetStudent struct {
	ID        string `json:"id"`
	Name 			string `json:"name"`
	Email     string `json:"email"`
}