package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type enrollmentResolver struct {
	enrollment *models.Enrollment
}

func (r *enrollmentResolver) ID() string {
	return r.enrollment.ID
}
func (r *enrollmentResolver) CourseID() string {
	return r.enrollment.CourseID
}
func (r *enrollmentResolver) UserID() string {
	return r.enrollment.UserID
}

func (r *enrollmentResolver) Course(ctx context.Context) *courseResolver {
	return &courseResolver{course: r.enrollment.Course}
}
func (r *enrollmentResolver) User(ctx context.Context) *userResolver {
	return &userResolver{user: r.enrollment.User}
}

func (r *enrollmentResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.enrollment.CreatedAt}, nil
}
func (r *enrollmentResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.enrollment.UpdatedAt}, nil
}
