package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Company : company
type Company struct {
	ID        string    `db:"id",json:"_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name",json:"name"`
	Courses   []*Course `json:"courses"`
}

func (c Company) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("Name is required")),
	)
}
