package resolvers

import (
	"encoding/json"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Error struct {
	Key     string
	Message string
}

type errorResolver struct {
	Error *Error
}

func (r *errorResolver) Key() *string {
	return &r.Error.Key
}

func (r *errorResolver) Message() *string {
	return &r.Error.Message
}

func resolveErrors(errors []*Error) *[]*errorResolver {
	resolvers := make([]*errorResolver, 0)
	for _, error := range errors {
		resolvers = append(resolvers, &errorResolver{Error: error})
	}
	return &resolvers
}

func mapErrors(err error) []*Error {
	errors := make([]*Error, 0)
	errorsmap := make(map[string]string, 0)
	b, _ := json.Marshal(err.(validation.Errors))
	if err := json.Unmarshal(b, &errorsmap); err != nil {
		fmt.Println("Error", err)
	}
	for f, v := range errorsmap {
		errors = append(errors, &Error{Key: f, Message: v})
	}
	return errors
}
