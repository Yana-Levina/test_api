package mysql

import (
	"context"
	"database/sql"
	_ "fmt"
	"github.com/sirupsen/logrus"
	"test/app"

	//_"github.com/gocraft/dbr"
	//_ "github.com/lib/pq"
	_ "github.com/labstack/echo/v4"
)

type mysqlPersonRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlPersonRepository(Conn *sql.DB) Person_repository {
	return &mysqlPersonRepository{Conn}
}

func (m *mysqlPersonRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []app.Person, err error) {
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

func (m *mysqlPersonRepository) GetByID(ctx context.Context, id int64) (err error) {

	query := `SELECT id, email, phone, first_name, last_name FROM person WHERE id = ?`
	//list, err := m.fetch(ctx, query, id)
	//if err != nil {
	//	return app.Person{}, err
	//}
	//
	//if len(list) > 0 {
	//	res = list[0]
	//} else {
	//	return res, app.Person
	//}
	//
	//return
}

func (_m *mysqlPersonRepository) Update(ctx context.Context, person *app.Person) (err error) {

	query := `UPDATE person SET email = ? , phone = ? , first_name = ?, last_name = ? WHERE id = ?`

}

func (m *mysqlPersonRepository) Delete(ctx context.Context, id int64) (err error) {

	query := `DELETE FROM person WHERE id = ?`
	//id, _ := strconv.Atoi(c.Param("id"))
	//delete(users, id)
	//return c.NoContent(http.StatusNoContent)
	return nil
}
