package resolvers

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type categoryResolver struct {
	category *models.Category
}

func (r *categoryResolver) ID(ctx context.Context) string {
	return r.category.ID
}
func (r *categoryResolver) Name(ctx context.Context) string {
	return r.category.Name
}
func (r *categoryResolver) ParentID(ctx context.Context) *string {
	return r.category.ParentID
}
func (r *categoryResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.category.CreatedAt}, nil
}
func (r *categoryResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.category.UpdatedAt}, nil
}
func (r *categoryResolver) Courses(ctx context.Context) *[]*courseResolver {
	l := make([]*courseResolver, len(r.category.Courses))
	for i := range l {
		l[i] = &courseResolver{course: r.category.Courses[i]}
	}
	return &l
}
