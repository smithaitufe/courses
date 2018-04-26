package resolvers

import (
	"context"
	"time"

	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

type CompanyInput struct {
	Name string
}

func (r *SchemaResolver) CreateCompany(ctx context.Context, args *struct {
	Input *CompanyInput
}) (*payloadResolver, error) {
	company := &models.Company{
		Name:      args.Input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := company.Validate(); err != nil {
		return &payloadResolver{ok: false, errors: mapErrors(err)}, nil
	}
	company, err := ctx.Value(ck.CompanyServiceKey).(*services.CompanyService).CreateCompany(company)
	if err != nil {
		return &payloadResolver{ok: false, errors: nil}, err
	}
	return &payloadResolver{ok: true, errors: nil}, nil

}
