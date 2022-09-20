package mysql

import (
	"context"
	"test/app"

	"github.com/stretchr/testify/mock"
)

type Person_repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Person_repository) GetByID(ctx context.Context, id int64) (app.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 app.Person
	if rf, ok := ret.Get(0).(func(context.Context, int64) app.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(app.Person)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ar
func (_m *Person_repository) Update(ctx context.Context, ar *app.Person) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *app.Person) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Person_repository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: _a0, _a1
//func (_m *Person_repository) Store(_a0 context.Context, _a1 *app.Person) error {
//	ret := _m.Called(_a0, _a1)
//
//	var r0 error
//	if rf, ok := ret.Get(0).(func(context.Context, *app.Person) error); ok {
//		r0 = rf(_a0, _a1)
//	} else {
//		r0 = ret.Error(0)
//	}
//
//	return r0
//}
