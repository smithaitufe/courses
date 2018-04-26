package resolvers

import (
	"context"
	"log"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/services"
)

type companyArgs struct {
	ID *string
}

func (r *SchemaResolver) Companies(ctx context.Context) (*[]*companyResolver, error) {
	companies, err := ctx.Value(ck.CompanyServiceKey).(*services.CompanyService).GetCompanies()
	if err != nil {
		log.Fatal(err)
	}
	resolvers := make([]*companyResolver, 0)
	for _, company := range companies {
		resolvers = append(resolvers, &companyResolver{company: company})
	}

	return &resolvers, nil
}

func (r *SchemaResolver) Company(ctx context.Context, args companyArgs) (*companyResolver, error) {
	company, err := ctx.Value(ck.CompanyServiceKey).(*services.CompanyService).GetCompany(args.ID)
	if err != nil {
		log.Fatal(err)
	}
	return &companyResolver{company: company}, nil
}
