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

func newCourseLoader() dataloader.BatchFunc {
	return courseLoader{}.loadBatch
}

type courseLoader struct{}

func (ldr courseLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		results = make([]*dataloader.Result, 0)
		wg      sync.WaitGroup
	)

	for i, key := range keys {
		wg.Add(1)
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			keyValue := key.String()
			course, err := ctx.Value(ck.CourseServiceKey).(*services.CourseService).GetCourse(&keyValue)
			results = append(results, &dataloader.Result{Data: course, Error: err})
		}(i, key)
	}

	wg.Wait()

	return results
}
func LoadCourses(ctx context.Context, keys []string) ([]*models.Course, error) {
	courses := make([]*models.Course, len(keys))

	ldr, err := extract(ctx, ck.CourseLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(keys))()

	if errs != nil {
		return nil, errs[0]
	}

	for _, d := range data {
		course, ok := d.(*models.Course)
		if !ok {
			return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", course, d)
		}
		courses = append(courses, course)
	}
	return courses, nil
}
func LoadCourse(ctx context.Context, key string) (*models.Course, error) {
	var course *models.Course

	ldr, err := extract(ctx, ck.CourseLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	course, ok := data.(*models.Course)
	if !ok {
		return nil, fmt.Errorf("Wrong type: the expected type is %T but got %T", course, data)
	}

	return course, nil
}
