package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type userResolver struct {
	user *models.User
}

func (r *userResolver) ID() string {
	return r.user.ID
}
func (r *userResolver) FirstName() string {
	return r.user.FirstName
}
func (r *userResolver) LastName() string {
	return r.user.LastName
}
func (r *userResolver) Email() string {
	return r.user.Email
}
func (r *userResolver) Country() string {
	return r.user.Country
}
func (r *userResolver) DialingCode() string {
	return r.user.DialingCode
}
func (r *userResolver) PhoneNumber() string {
	return r.user.PhoneNumber
}
func (r *userResolver) Roles(ctx context.Context) *[]*roleResolver {
	l := make([]*roleResolver, 0, len(r.user.Roles))
	for i := range l {
		l[i] = &roleResolver{role: r.user.Roles[i]}
	}
	return &l
}
func (r *userResolver) Enrollments(ctx context.Context) *[]*enrollmentResolver {
	l := make([]*enrollmentResolver, 0, len(r.user.Enrollments))
	for i := range l {
		l[i] = &enrollmentResolver{enrollment: r.user.Enrollments[i]}
	}
	return &l
}
func (r *userResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.user.CreatedAt}, nil
}
func (r *userResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.user.UpdatedAt}, nil
}
