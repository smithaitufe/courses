package models

import (
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              string    `db:"id",json:"_id"`
	FirstName       string    `db:"first_name",json:"firstName"`
	LastName        string    `db:"last_name",json:"lastName"`
	Email           string    `db:"email",json:"email"`
	Country         string    `db:"country",json:"country"`
	DialingCode     string    `db:"dialing_code",json:"dialingCode"`
	PhoneNumber     string    `db:"phone_number",json:"phoneNumber"`
	Password        string    `db:"password",json:"password"`
	ConfirmPassword string    `db:"-",json:"confirmPassword"`
	CreatedAt       time.Time `db:"created_at",json:"createdAt"`
	UpdatedAt       time.Time `db:"updated_at",json:"updatedAt"`
	Roles           []*Role
	Enrollments     []*Enrollment
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.FirstName, validation.Required.Error("FirstName is required")),
		validation.Field(&user.LastName, validation.Required.Error("LastName is required")),
		validation.Field(&user.Email, validation.Required.Error("Email is required")),
		validation.Field(&user.PhoneNumber, validation.Required.Error("PhoneNumber is required")),
		validation.Field(&user.Password, validation.Required.Error("Password is required")),
	)
}
func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
