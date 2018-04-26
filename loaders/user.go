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

func newUserLoader() dataloader.BatchFunc {
	return userLoader{}.loadBatch
}

type userLoader struct{}

func (ldr userLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			user, err := ctx.Value("userService").(*services.UserService).GetUser(&keyValue)
			results = append(results, &dataloader.Result{Data: user, Error: err})
		}(i, key)
	}

	wg.Wait()

	return results
}
func LoadUsers(ctx context.Context, keys []string) ([]*models.User, error) {
	users := make([]*models.User, len(keys))

	ldr, err := extract(ctx, ck.UserLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(keys))()

	if errs != nil {
		return nil, errs[0]
	}

	for _, d := range data {
		user, ok := d.(*models.User)
		if !ok {
			return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", user, d)
		}
		users = append(users, user)
	}
	return users, nil
}
func LoadUser(ctx context.Context, key string) (*models.User, error) {
	var user *models.User

	ldr, err := extract(ctx, ck.UserLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	user, ok := data.(*models.User)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", user, data)
	}

	return user, nil
}
