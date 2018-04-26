package loaders

import (
	"context"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader"
	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/models"
	"github.com/smithaitufe/courses/services"
)

func newCompanyLoader() dataloader.BatchFunc {
	return companyLoader{}.loadBatch
}

type companyLoader struct{}

func (ldr companyLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			company, err := ctx.Value(ck.CompanyServiceKey).(*services.CompanyService).GetCompany(&keyValue)
			results = append(results, &dataloader.Result{Data: company, Error: err})
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadCompany(ctx context.Context, key string) (*models.Company, error) {
	var company *models.Company

	ldr, err := extract(ctx, ck.CompanyLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	company, ok := data.(*models.Company)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", company, data)
	}

	return company, nil
}
