package dto

type UserCreatedDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Identification string `json:"identification"`
	Status         string `json:"status"`
} // @name userResponse
