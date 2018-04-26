package resolvers

import (
	"context"

	graphql "github.com/neelance/graphql-go"
	"github.com/smithaitufe/courses/models"
)

type roleResolver struct {
	role *models.Role
}

func (r *roleResolver) ID() string {
	return r.role.ID
}
func (r *roleResolver) Name() string {
	return r.role.Name
}
func (r *roleResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.role.CreatedAt}, nil
}
func (r *roleResolver) UpdatedAt(ctx context.Context) (graphql.Time, error) {
	return graphql.Time{Time: r.role.UpdatedAt}, nil
}
