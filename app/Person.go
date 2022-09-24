package app

import "context"

type Person struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PersonRepository interface {
	Create(context.Context, *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, person *Person) error
	Delete(ctx context.Context, id int64) error
}

type PersonUsecase interface {
	Create(context.Context, *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, person *Person) error
	Delete(ctx context.Context, id int64) error
}
