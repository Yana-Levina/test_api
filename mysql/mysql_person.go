package mysql

import (
	"context"
	"database/sql"
	"errors"
	_ "fmt"
	"github.com/sirupsen/logrus"
	"test/app"

	//_"github.com/gocraft/dbr"
	//_ "github.com/lib/pq"
	_ "github.com/labstack/echo/v4"
)

type PersonRepository struct {
	Conn *sql.DB
}

func NewPersonRepository(Conn *sql.DB) app.PersonRepository {
	return &PersonRepository{Conn}
}

func (m *PersonRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []app.Person, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]app.Person, 0)
	for rows.Next() {
		t := app.Person{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Phone,
			&t.FirstName,
			&t.LastName,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *PersonRepository) Create(ctx context.Context, person *app.Person) error {
	//TODO implement me
	panic("implement me")
}

func (m *PersonRepository) GetAll(ctx context.Context) (res []app.Person, err error) {
	query := `SELECT id, email, phone, first_name, last_name FROM public.person`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *PersonRepository) GetByID(ctx context.Context, id int64) (res app.Person, err error) {

	query := `SELECT id, email, phone, first_name, last_name FROM public.person WHERE id = $1`
	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return app.Person{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, errors.New("your requested Item is not found")
	}

	return
}

func (_m *PersonRepository) Update(ctx context.Context, person *app.Person) (err error) {

	//query := `UPDATE person SET email = $1 , phone = $1 , first_name = $1, last_name = ? WHERE id = $1`
	return nil

}

func (m *PersonRepository) Delete(ctx context.Context, id int64) (err error) {

	//query := `DELETE FROM person WHERE id = ?`
	//id, _ := strconv.Atoi(c.Param("id"))
	//delete(users, id)
	//return c.NoContent(http.StatusNoContent)
	return nil
}
