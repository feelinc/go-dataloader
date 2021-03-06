//go:generate go run github.com/feelinc/go-dataloader UserLoader string *github.com/feelinc/go-dataloader/example.User

package example

import (
	"time"
)

// User is some kind of database backed model
type User struct {
	ID     string
	Name   string
	Params []string
}

// NewLoader will collect user requests for 2 milliseconds and send them as a single batch to the fetch func
// normally fetch would be a database call.
func NewLoader() *UserLoader {
	return &UserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string, params ...[]interface{}) ([]*User, []error) {
			users := make([]*User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = &User{ID: key, Name: "user " + key}
			}
			return users, errors
		},
	}
}
