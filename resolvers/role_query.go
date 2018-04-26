package resolvers

import (
	"context"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type rolesArgs struct {
}
type roleArgs struct {
	ID string
}

func (r *SchemaResolver) Roles(ctx context.Context) (*[]*roleResolver, error) {
	roles := make([]*models.Role, 0)

	roles, err := ctx.Value(ck.RoleServiceKey).(*services.RoleService).GetRoles()
	if err != nil {
		return nil, err
	}
	resolvers := make([]*roleResolver, 0)
	for _, role := range roles {
		resolvers = append(resolvers, &roleResolver{role: role})
	}
	return &resolvers, nil
}
func (r *SchemaResolver) Role(ctx context.Context, args roleArgs) (*roleResolver, error) {
	role, err := ctx.Value(ck.RoleServiceKey).(*services.RoleService).GetRole(args.ID)
	if err != nil {
		return nil, err
	}
	return &roleResolver{role: role}, nil
}
