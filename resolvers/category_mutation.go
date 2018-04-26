package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type CategoryInput struct {
	Name     string
	ParentID *string
}

func (r *SchemaResolver) CreateCategory(ctx context.Context, args *struct {
	Input *CategoryInput
}) (*payloadResolver, error) {
	category := &models.Category{
		Name:      args.Input.Name,
		ParentID:  args.Input.ParentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := category.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	category, err := ctx.Value(ck.CategoryServiceKey).(*services.CategoryService).CreateCategory(category)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil

}
