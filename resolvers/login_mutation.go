package resolvers

import (
	"context"

	ce "github.com/smithaitufe/courses/errors"
	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/services"
)

type loginInput struct {
	Username string
	Password string
}

type loginPayload struct {
	ok           bool
	errors       []*Error
	token        *string
	refreshToken *string
}

func (r *loginPayload) Ok() bool {
	return r.ok
}
func (r *loginPayload) Errors() *[]*errorResolver {
	return resolveErrors(r.errors)
}
func (r *loginPayload) Token() *string {
	return r.token
}
func (r *loginPayload) RefreshToken() *string {
	return r.refreshToken
}

func (r *SchemaResolver) Login(ctx context.Context, args *struct {
	Input *loginInput
}) (*loginPayload, error) {
	loginErr := &Error{Key: "login", Message: "Invalid login credentials. Check your username or password"}

	user, err := ctx.Value(ck.UserServiceKey).(*services.UserService).FindUserByEmail(args.Input.Username)
	ce.LogOnError("Could not fetch user", err)

	if user == nil {
		user, err := ctx.Value(ck.UserServiceKey).(*services.UserService).FindUserByPhoneNumber(args.Input.Username)
		ce.LogOnError("Could not fetch user", err)
		if user == nil {
			return &loginPayload{ok: false, errors: []*Error{loginErr}}, nil
		}
	}
	if user.ComparePassword(args.Input.Password) {
		token, err := ctx.Value(ck.AuthServiceKey).(*services.AuthService).GenerateToken(user)
		ce.LogOnError("Could not generate token", err)

		refreshToken, err := ctx.Value(ck.AuthServiceKey).(*services.AuthService).GenerateRefreshToken(user)
		ce.LogOnError("Could not generate refresh-token", err)
		return &loginPayload{ok: true, errors: nil, token: token, refreshToken: refreshToken}, nil
	}

	return &loginPayload{ok: false, errors: []*Error{loginErr}}, nil
}
