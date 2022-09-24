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

func (p *PersonUsecase) GetAll(c context.Context) (res []app.Person, err error) {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	res, err = p.personRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	//res, err = p.fillAuthorDetails(ctx, res)
	//if err != nil {
	//	nextCursor = ""
	//}
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

func (p PersonUsecase) Update(ctx context.Context, person *app.Person) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonUsecase) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
