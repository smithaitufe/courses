package resolvers

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type companyResolver struct {
	company *models.Company
}

func (r *companyResolver) ID(ctx context.Context) string {
	return r.company.ID
}
func (r *companyResolver) Name(ctx context.Context) string {
	return r.company.Name
}
func (r *companyResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.company.CreatedAt}, nil
}
func (r *companyResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.company.UpdatedAt}, nil
}
func (r *companyResolver) Courses(ctx context.Context) *[]*courseResolver {
	l := make([]*courseResolver, 0, len(r.company.Courses))
	for i := range l {
		l[i] = &courseResolver{course: r.company.Courses[i]}
	}
	return &l
}
