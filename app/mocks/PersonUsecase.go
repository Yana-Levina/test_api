package mocks

import (
	"context"
	mock "github.com/stretchr/testify/mock"
	"test/app"
)

type PersonUsecase struct {
	mock.Mock
}

func (_m *PersonUsecase) Update(ctx context.Context, ar *app.Person) error {
	//ret := _m.Called(ctx, ar)
	//
	//var r0 error
	//if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
	//	r0 = rf(ctx, ar)
	//} else {
	//	r0 = ret.Error(0)
	//}

	return nil
}

func (_m *PersonUsecase) Delete(ctx context.Context, id int64) error {
	//ret := _m.Called(ctx, id)
	//
	//var r0 error
	//if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
	//	r0 = rf(ctx, id)
	//} else {
	//	r0 = ret.Error(0)
	//}

	return nil
}
