package postgres

import (
	_ "context"
	_ "database/sql"
	_ "test/app"

	_ "github.com/labstack/echo/v4"
)

//func Create() {
//
//	query := `INSERT person SET email = ? , phone = ? , first_name = ?, last_name = ?`;
//
//}
//
//func GetByID(ctx context.Context, id int64)  {
//
//	query := `SELECT id, email, phone, first_name, last_name FROM person WHERE id = ?`;
//
//}
//
//func Update(ctx context.Context, person *app.Person){
//
//	query := `UPDATE person SET email = ? , phone = ? , first_name = ?, last_name = ? WHERE id = ?`;
//
//}
//
//func Delete(ctx context.Context, id int64)  error {
//
//	query := `DELETE FROM person WHERE id = ?`;
//	//id, _ := strconv.Atoi(c.Param("id"))
//	//delete(users, id)
//	//return c.NoContent(http.StatusNoContent)
//	return nil;
//}
