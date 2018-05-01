package resolvers

import (
	"context"
	"log"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/loaders"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type coursesArgs struct {
	Title *string
}
type courseArgs struct {
	ID *string
}

func (r *SchemaResolver) Courses(ctx context.Context, args coursesArgs) (*[]*courseResolver, error) {
	courses := make([]*models.Course, 0)

	courses, err := ctx.Value(ck.CourseServiceKey).(*services.CourseService).GetCourses(args.Title)
	if err != nil {
		log.Fatal(err)
	}
	resolvers := make([]*courseResolver, 0)
	for _, course := range courses {
		if company, err := loaders.LoadCompany(ctx, course.CompanyID); err == nil {
			course.Company = company
		}
		if category, err := loaders.LoadCategory(ctx, course.CategoryID); err == nil {
			course.Category = category
		}
		resolvers = append(resolvers, &courseResolver{course: course})
	}

	return &resolvers, nil
}
func (r *SchemaResolver) Course(ctx context.Context, args courseArgs) (*courseResolver, error) {
	course, err := ctx.Value(ck.CourseServiceKey).(*services.CourseService).GetCourse(args.ID)
	if company, err := loaders.LoadCompany(ctx, course.CompanyID); err == nil {
		course.Company = company
	}
	if category, err := loaders.LoadCategory(ctx, course.CategoryID); err == nil {
		course.Category = category
	}
	if err != nil {
		return nil, err
	}
	return &courseResolver{course: course}, nil
}
