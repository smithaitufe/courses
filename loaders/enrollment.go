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

func newEnrollmentLoader() dataloader.BatchFunc {
	return enrollmentLoader{}.loadBatch
}

type enrollmentLoader struct{}

func (ldr enrollmentLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			user, err := ctx.Value(ck.EnrollmentServiceKey).(*services.EnrollmentService).GetEnrollment(keyValue)
			results = append(results, &dataloader.Result{Data: user, Error: err})
		}(i, key)
	}

	wg.Wait()

	return results
}
func LoadEnrollments(ctx context.Context, keys []string) ([]*models.Enrollment, error) {
	enrollments := make([]*models.Enrollment, len(keys))

	ldr, err := extract(ctx, ck.EnrollmentLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(keys))()

	if errs != nil {
		return nil, errs[0]
	}

	for _, d := range data {
		enrollment, ok := d.(*models.Enrollment)
		if !ok {
			return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", enrollment, d)
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
func LoadEnrollment(ctx context.Context, key string) (*models.Enrollment, error) {
	var enrollment *models.Enrollment

	ldr, err := extract(ctx, ck.EnrollmentLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	enrollment, ok := data.(*models.Enrollment)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", enrollment, data)
	}

	return enrollment, nil
}
