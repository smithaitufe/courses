package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type EnrollmentInput struct {
	UserID   string
	CourseID string
}

func (r *SchemaResolver) CreateEnrollment(ctx context.Context, args *struct {
	Input *EnrollmentInput
}) (*payloadResolver, error) {

	enrollment := &models.Enrollment{
		UserID:    args.Input.UserID,
		CourseID:  args.Input.CourseID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := enrollment.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	enrollment, err := ctx.Value(ck.EnrollmentServiceKey).(*services.EnrollmentService).CreateEnrollment(enrollment)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil
}
