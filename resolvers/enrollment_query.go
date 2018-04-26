package resolvers

import (
	"context"
	"log"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/loaders"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type enrollmentsArgs struct {
	UserID   *string
	CourseID *string
}
type enrollmentArgs struct {
	ID string
}

func (r *SchemaResolver) Enrollments(ctx context.Context, args enrollmentsArgs) (*[]*enrollmentResolver, error) {
	enrollments := make([]*models.Enrollment, 0)

	enrollments, err := ctx.Value(ck.EnrollmentServiceKey).(*services.EnrollmentService).GetEnrollments()
	if err != nil {
		log.Fatal(err)
	}
	resolvers := make([]*enrollmentResolver, 0)
	for _, enrollment := range enrollments {
		if user, err := loaders.LoadUser(ctx, enrollment.UserID); err == nil {
			enrollment.User = user
		}
		if course, err := loaders.LoadCourse(ctx, enrollment.CourseID); err == nil {
			enrollment.Course = course
		}
		resolvers = append(resolvers, &enrollmentResolver{enrollment: enrollment})
	}
	return &resolvers, nil
}
func (r *SchemaResolver) Enrollment(ctx context.Context, args enrollmentArgs) (*enrollmentResolver, error) {
	enrollment, err := ctx.Value(ck.EnrollmentServiceKey).(*services.EnrollmentService).GetEnrollment(args.ID)
	if user, err := loaders.LoadUser(ctx, enrollment.UserID); err == nil {
		enrollment.User = user
	}
	if course, err := loaders.LoadUser(ctx, enrollment.UserID); err == nil {
		enrollment.User = course
	}
	if err != nil {
		return nil, err
	}
	return &enrollmentResolver{enrollment: enrollment}, nil
}
