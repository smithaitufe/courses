package resolvers

import "context"

type SchemaResolver struct{}

type payloadResolver struct {
	ok     bool
	errors []*Error
}

func (r *payloadResolver) Ok(ctx context.Context) bool {
	return r.ok
}
func (r *payloadResolver) Errors(ctx context.Context) *[]*errorResolver {
	return resolveErrors(r.errors)
}
