package models

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
)

// Course : course
type Course struct {
	ID          string    `db:"id",json:"_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Code        string    `db:"code",json:"code"`
	Title       string    `db:"title",json:"title"`
	Amount      float64   `db:"amount",json:"amount"`
	Hours       int32     `db:"hours",json:"hours"`
	CategoryID  string    `db:"category_id",json:"category_id"`
	Category    *Category `json:"category"`
	CompanyID   string    `db:"company_id",json:"company_id"`
	Company     *Company  `json:"company"`
	Overview    string    `db:"overview",json:"overview"`
	Description string    `db:"description",json:"description"`
}

func (c Course) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Code, validation.Required.Error("Code is required")),
		validation.Field(&c.Title, validation.Required.Error("Title is required")),
		validation.Field(&c.Hours, validation.Required.Error("Hours is required")),
		validation.Field(&c.CategoryID, validation.Required.Error("Category is required")),
		validation.Field(&c.CompanyID, validation.Required.Error("Company is required")),
	)
}
