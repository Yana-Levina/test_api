package postgres

import (
	"test/app"

	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PersonRepository struct {
	Conn *sql.DB
}

func NewPersonRepository(Conn *sql.DB) app.PersonRepository {
	return &PersonRepository{Conn}
}

func (pg *PersonRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []app.Person, err error) {
	rows, err := pg.Conn.QueryContext(ctx, query, args...)
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

func (pg *PersonRepository) Create(ctx context.Context, person *app.Person) (err error) {
	query := `INSERT INTO public.person(email, phone, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`

	stmt, err := pg.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	var id int
	err = stmt.QueryRowContext(ctx, person.Email, person.Phone, person.FirstName, person.LastName).Scan(&id)

	if err != nil {
		return
	}
	person.ID = int64(id)
	return
}

func (pg *PersonRepository) GetAll(ctx context.Context) (res []app.Person, err error) {
	query := `SELECT id, email, phone, first_name, last_name FROM public.person`

	res, err = pg.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (pg *PersonRepository) GetByID(ctx context.Context, id int64) (res app.Person, err error) {
	query := `SELECT id, email, phone, first_name, last_name FROM public.person WHERE id = $1`
	list, err := pg.fetch(ctx, query, id)
	if err != nil {
		return app.Person{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, errors.New("Your requested person is not found")
	}

	return
}

func (pg *PersonRepository) Update(ctx context.Context, person *app.Person, id int64) (err error) {
	query := `UPDATE public.person SET email = $1 , phone = $2 , first_name = $3, last_name = $4 WHERE id = $5`
	stmt, err := pg.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, person.Email, person.Phone, person.FirstName, person.LastName, id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}

func (pg *PersonRepository) Delete(ctx context.Context, id int64) (err error) {

	query := `DELETE FROM person WHERE id = $1`
	stmt, err := pg.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
