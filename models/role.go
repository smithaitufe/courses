package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Role struct {
	ID        string    `db:"id",json:"_id"`
	Name      string    `db:"name",json:"name"`
	CreatedAt time.Time `db:"created_at",json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at",json:"updatedAt"`
}

func (role Role) Validate() error {
	return validation.ValidateStruct(&role,
		validation.Field(&role.Name, validation.Required.Error("Name is required")),
	)
}
