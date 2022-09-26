package logic

import (
	"context"
	"errors"
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

func (p PersonUsecase) Create(c context.Context, person *app.Person) (err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	err = p.personRepo.Create(ctx, person)
	return
}

func (p *PersonUsecase) GetAll(c context.Context) (res []app.Person, err error) {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	res, err = p.personRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (p PersonUsecase) GetByID(c context.Context, id int64) (res app.Person, err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	res, err = p.personRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}

// iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii
func (p PersonUsecase) Update(c context.Context, person *app.Person) (err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.personRepo.Update(ctx, person)
}

func (p PersonUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	existedArticle, err := p.personRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (app.Person{}) {
		return errors.New("person is not found")
	}
	return p.personRepo.Delete(ctx, id)
}
