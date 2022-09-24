package logic

import (
	"context"
	"test/app"
	"time"
)

type PersonUsecase struct {
	personRepo     app.PersonRepository
	contextTimeout time.Duration
}

func NewPersonUsecase(a app.PersonRepository, timeout time.Duration) app.PersonUsecase {
	return &PersonUsecase{
		personRepo:     a,
		contextTimeout: timeout,
	}
}

func (p PersonUsecase) Create(ctx context.Context, person *app.Person) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonUsecase) GetAll(ctx context.Context, cursor string, num int64) ([]app.Person, string, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonUsecase) GetByID(ctx context.Context, id int64) (app.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonUsecase) Update(ctx context.Context, person *app.Person) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonUsecase) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
