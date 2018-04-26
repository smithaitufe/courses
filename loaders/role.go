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

func newRoleLoader() dataloader.BatchFunc {
	return roleLoader{}.loadBatch
}

type roleLoader struct{}

func (ldr roleLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			user, err := ctx.Value(ck.RoleServiceKey).(*services.RoleService).GetRole(keyValue)
			results = append(results, &dataloader.Result{Data: user, Error: err})
		}(i, key)
	}

	wg.Wait()

	return results
}
func LoadRoles(ctx context.Context, keys []string) ([]*models.Role, error) {
	roles := make([]*models.Role, len(keys))

	ldr, err := extract(ctx, ck.RoleLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(keys))()

	if errs != nil {
		return nil, errs[0]
	}

	for _, d := range data {
		role, ok := d.(*models.Role)
		if !ok {
			return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", role, d)
		}
		roles = append(roles, role)
	}
	return roles, nil
}
func LoadRole(ctx context.Context, key string) (*models.Role, error) {
	var role *models.Role

	ldr, err := extract(ctx, ck.RoleLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	role, ok := data.(*models.Role)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", role, data)
	}

	return role, nil
}
