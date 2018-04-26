package resolvers

import (
	"context"
	"log"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/services"
)

type categoryArgs struct {
	ID *string
}

func (r *SchemaResolver) Categories(ctx context.Context) (*[]*categoryResolver, error) {
	categories, err := ctx.Value(ck.CategoryServiceKey).(*services.CategoryService).GetCategories()
	if err != nil {
		return nil, err
	}
	resolvers := make([]*categoryResolver, 0)
	for _, category := range categories {
		resolvers = append(resolvers, &categoryResolver{category: category})
	}

	return &resolvers, nil
}
func (r *SchemaResolver) Category(ctx context.Context, args categoryArgs) (*categoryResolver, error) {
	category, err := ctx.Value(ck.CategoryServiceKey).(*services.CategoryService).GetCategory(args.ID)
	if err != nil {
		log.Fatal(err)
	}

	return &categoryResolver{category: category}, nil
}
