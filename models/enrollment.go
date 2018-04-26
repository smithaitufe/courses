package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Enrollment struct {
	ID        string `db:"id",json:"_id"`
	UserID    string `db:"user_id",json:"userId"`
	User      *User
	CreatedAt time.Time `db:"created_at",json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at",json:"updatedAt"`
	CourseID  string    `db:"course_id",json:"courseId"`
	Course    *Course
}

func (enrollment Enrollment) Validate() error {
	return validation.ValidateStruct(&enrollment,
		validation.Field(&enrollment.UserID, validation.Required.Error("User ID is required")),
		validation.Field(&enrollment.CourseID, validation.Required.Error("Course ID is required")),
	)
}
