package postgres

import (
	"context"
	"database/sql"
	_ "database/sql"
	_ "fmt"
	"test/app"
	//_"github.com/gocraft/dbr"

	//_ "github.com/lib/pq"
	_ "github.com/labstack/echo/v4"
)

type postgresPersonRepo struct {
	DB *sql.DB
}

func createPerson() {

	query := `INSERT person SET email = ? , phone = ? , first_name = ?, last_name = ?`

}

func GetByID(ctx context.Context, id int64) {

	query := `SELECT id, email, phone, first_name, last_name FROM person WHERE id = ?`

}

func Update(ctx context.Context, person *app.Person) {

	query := `UPDATE person SET email = ? , phone = ? , first_name = ?, last_name = ? WHERE id = ?`

}

func Delete(ctx context.Context, id int64) error {

	query := `DELETE FROM person WHERE id = ?`
	//id, _ := strconv.Atoi(c.Param("id"))
	//delete(users, id)
	//return c.NoContent(http.StatusNoContent)
	return nil
}
