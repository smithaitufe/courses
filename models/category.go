package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Semester : semester
type Category struct {
	ID        string    `db:"id",json:"_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name",json:"name"`
	ParentID  *string   `db:"parent_id",json:"parent_id"`
	Courses   []*Course `json:"courses"`
}

func (c Category) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("Name is required")),
	)
}
