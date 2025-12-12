package domain

type Contact struct {
	Id      int    `json:"id" db:"id"`
	Email   string `json:"email" db:"email" validate:"required,email"`
	Name    string `json:"name" db:"name" validate:"required"`
	Surname string `json:"surname" db:"surname" validate:"required"`
	Company string `json:"company" db:"company"`
	Message string `json:"message" db:"message" validate:"required"`
}
