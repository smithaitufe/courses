package resolvers

import (
	"context"
	"log"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/services"
)

type userArgs struct {
	ID *string
}

func (r *SchemaResolver) Users(ctx context.Context) (*[]*userResolver, error) {
	users, err := ctx.Value(ck.UserServiceKey).(*services.UserService).GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	resolvers := make([]*userResolver, 0)
	for _, user := range users {
		if roles, err := ctx.Value(ck.RoleServiceKey).(*services.RoleService).FindRolesByUserId(user.ID); err == nil {
			user.Roles = roles
		}
		if enrollments, err := ctx.Value(ck.EnrollmentServiceKey).(*services.EnrollmentService).FindEnrollmentsByUserId(user.ID); err == nil {
			user.Enrollments = enrollments
		}
		resolvers = append(resolvers, &userResolver{user: user})
	}

	return &resolvers, nil
}

func (r *SchemaResolver) User(ctx context.Context, args userArgs) (*userResolver, error) {
	user, err := ctx.Value(ck.UserServiceKey).(*services.UserService).GetUser(args.ID)
	if err != nil {
		log.Fatal(err)
	}
	return &userResolver{user: user}, nil
}
