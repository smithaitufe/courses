package resolvers

import "context"

type SchemaResolver struct{}

// func parseTime(value string) (graphql.Time, error) {
// 	if value == "" {
// 		return nil, nil
// 	}
//
// 	t, err := time.Parse(time.RFC3339, value)
// 	return graphql.Time{Time: t}, err
// }
func sPtr(s string) *string { return &s }

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
