package resolvers

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type courseResolver struct {
	course *models.Course
}

func (r *courseResolver) ID(ctx context.Context) string {
	return r.course.ID
}
func (r *courseResolver) CategoryID(ctx context.Context) string {
	return r.course.CategoryID
}
func (r *courseResolver) CompanyID(ctx context.Context) string {
	return r.course.CompanyID
}
func (r *courseResolver) Code(ctx context.Context) string {
	return r.course.Code
}
func (r *courseResolver) Title(ctx context.Context) string {
	return r.course.Title
}
func (r *courseResolver) Overview(ctx context.Context) string {
	return r.course.Overview
}
func (r *courseResolver) Description(ctx context.Context) string {
	return r.course.Description
}
func (r *courseResolver) Amount(ctx context.Context) float64 {
	return r.course.Amount
}
func (r *courseResolver) Hours(ctx context.Context) int32 {
	return r.course.Hours
}
func (r *courseResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.course.CreatedAt}, nil
}
func (r *courseResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.course.UpdatedAt}, nil
}
func (r *courseResolver) Company(ctx context.Context) *companyResolver {
	return &companyResolver{company: r.course.Company}
}
func (r *courseResolver) Category(ctx context.Context) *categoryResolver {
	return &categoryResolver{category: r.course.Category}
}
