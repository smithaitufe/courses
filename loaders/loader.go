package loaders

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	ck "github.com/smithaitufe/courses/keys"
)

// type key string

type LoaderCollection struct {
	lookup map[ck.Key]dataloader.BatchFunc
}

// const (
// 	categoryLoaderKey   key = "category"
// 	companyLoaderKey    key = "company"
// 	courseLoaderKey     key = "course"
// 	userLoaderKey       key = "user"
// 	enrollmentLoaderKey key = "enrollment"
// 	roleLoaderKey       key = "role"
// )

func NewLoaderCollection() LoaderCollection {
	return LoaderCollection{
		lookup: map[ck.Key]dataloader.BatchFunc{
			ck.CategoryLoaderKey:   newCategoryLoader(),
			ck.CompanyLoaderKey:    newCompanyLoader(),
			ck.CourseLoaderKey:     newCourseLoader(),
			ck.UserLoaderKey:       newUserLoader(),
			ck.RoleLoaderKey:       newRoleLoader(),
			ck.EnrollmentLoaderKey: newEnrollmentLoader(),
		},
	}
}

func (c LoaderCollection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.lookup {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn))
	}
	return ctx
}

func extract(ctx context.Context, k ck.Key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("Cannot load %s loader on the request context", k)
	}

	return ldr, nil
}
