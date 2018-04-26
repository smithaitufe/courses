package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type RoleInput struct {
	Name string
}

func (r *SchemaResolver) CreateRole(ctx context.Context, args *struct {
	Input *RoleInput
}) (*payloadResolver, error) {

	role := &models.Role{
		Name:      args.Input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := role.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	role, err := ctx.Value(ck.RoleServiceKey).(*services.RoleService).CreateRole(role)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil
}
