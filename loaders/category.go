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

func newCategoryLoader() dataloader.BatchFunc {
	return categoryLoader{}.loadBatch
}

type categoryLoader struct{}

func (ldr categoryLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for _, key := range keys {
		wg.Add(1)
		go func(key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			category, err := ctx.Value(ck.CategoryServiceKey).(*services.CategoryService).GetCategory(&keyValue)
			results = append(results, &dataloader.Result{Data: category, Error: err})
		}(key)
	}

	wg.Wait()

	return results
}

func LoadCategory(ctx context.Context, key string) (*models.Category, error) {
	var category *models.Category

	ldr, err := extract(ctx, ck.CategoryLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	category, ok := data.(*models.Category)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", category, data)
	}

	return category, nil
}
