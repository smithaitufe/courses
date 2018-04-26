package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type CourseInput struct {
	Title      string
	Code       string
	Hours      int32
	Amount     float64
	CategoryID string
	CompanyID  string
}

func (r *SchemaResolver) CreateCourse(ctx context.Context, args *struct {
	Input *CourseInput
}) (*payloadResolver, error) {

	course := &models.Course{
		Title:      args.Input.Title,
		Code:       args.Input.Code,
		Hours:      args.Input.Hours,
		Amount:     args.Input.Amount,
		CategoryID: args.Input.CategoryID,
		CompanyID:  args.Input.CompanyID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := course.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	course, err := ctx.Value(ck.CourseServiceKey).(*services.CourseService).CreateCourse(course)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil
}
