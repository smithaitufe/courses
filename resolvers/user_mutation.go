package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type UserInput struct {
	FirstName       string
	LastName        string
	Email           string
	PhoneNumber     string
	Country         string
	DialingCode     string
	Password        string
	ConfirmPassword string
}

func (r *SchemaResolver) CreateUser(ctx context.Context, args *struct {
	Input *UserInput
}) (*payloadResolver, error) {
	user := &models.User{
		FirstName:       args.Input.FirstName,
		LastName:        args.Input.LastName,
		Email:           args.Input.Email,
		PhoneNumber:     args.Input.PhoneNumber,
		DialingCode:     args.Input.DialingCode,
		Country:         args.Input.Country,
		Password:        args.Input.Password,
		ConfirmPassword: args.Input.ConfirmPassword,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := user.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	user, err := ctx.Value(ck.UserServiceKey).(*services.UserService).CreateUser(user)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil

}
