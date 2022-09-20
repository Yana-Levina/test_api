package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"
	"test/app"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PersonHandler struct {
	PUsecase app.PersonUsecase
}

func NewPersonHandler(e *echo.Echo, us app.PersonUsecase) {
	handler := &PersonHandler{
		PUsecase: us,
	}
	e.GET("/persons", handler.Fetch)
	e.POST("/persons", handler.Store)
	e.GET("/persons/:id", handler.GetByID)
	e.DELETE("/persons/:id", handler.Delete)
}

func (a *PersonHandler) Fetch(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listAr, nextCursor, err := a.PUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *PersonHandler) Store(c echo.Context) error {

	return nil
}

func (a *PersonHandler) GetByID(c echo.Context) error {

	return nil
}

func (a *PersonHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//return c.JSON(http.StatusNotFound, app.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.PUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	//if err == nil {
	return http.StatusOK
	//}

	//logrus.Error(err)
	//switch err {
	//case app.ErrInternalServerError:
	//	return http.StatusInternalServerError
	//case app.ErrNotFound:
	//	return http.StatusNotFound
	//case app.ErrConflict:
	//	return http.StatusConflict
	//default:
	//	return http.StatusInternalServerError
	//}
}
