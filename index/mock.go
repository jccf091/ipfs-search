package index

import (
	"context"
	"reflect"
)

type mockResult struct {
	references []string
}

type mockIndex struct {
	ID     string
	Fields []string
	Result mockResult
	Found  bool
	Error  error
}

func (m *mockIndex) Get(ctx context.Context, id string, dst interface{}, fields ...string) (bool, error) {
	m.ID = id
	m.Fields = fields

	// Set result
	v := reflect.ValueOf(dst).Elem()
	v.Set(reflect.ValueOf(m.Result))

	return m.Found, m.Error
}