package app

import "context"

type Person struct {
	ID        int64  `json:"id" form:"id" xml:"id" query:"id"`
	Email     string `json:"email" form:"email" xml:"email" query:"email"`
	Phone     string `json:"phone" form:"phone" xml:"phone" query:"phone"`
	FirstName string `json:"first_name" form:"first_name" xml:"first_name" query:"first_name"`
	LastName  string `json:"last_name" form:"last_name" xml:"last_name" query:"last_name"`
}

type PersonRepository interface {
	Create(context.Context, *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, person *Person, id int64) error
	Delete(ctx context.Context, id int64) error
}

type PersonUsecase interface {
	Create(context.Context, *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, person *Person, id int64) error
	Delete(ctx context.Context, id int64) error
}
