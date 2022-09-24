package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"test/app"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PersonHandler struct {
	PUseCase app.PersonUsecase
}

func NewPersonHandler(e *echo.Echo, pu app.PersonUsecase) {
	handler := &PersonHandler{
		PUseCase: pu,
	}
	e.POST("/persons", handler.Create)
	e.GET("/persons", handler.GetAll)
	e.GET("/persons/:id", handler.GetByID)
	e.PUT("/persons/:id", handler.Update)
	e.DELETE("/persons/:id", handler.Delete)
}

func (a *PersonHandler) Create(c echo.Context) error {
	return nil
}

func (a *PersonHandler) GetAll(c echo.Context) error {
	//numS := c.QueryParam("num")
	//num, _ := strconv.Atoi(numS)
	//cursor := c.QueryParam("cursor")
	//ctx := c.Request().Context()
	//
	//listAr, nextCursor, err := a.PLogic.Fetch(ctx, cursor, int64(num))
	//if err != nil {
	//	//return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	//}
	//
	//c.Response().Header().Set(`X-Cursor`, nextCursor)
	//return c.JSON(http.StatusOK, listAr)
	return c.String(http.StatusOK, "you tru get all")
}

func (a *PersonHandler) GetByID(c echo.Context) error {
	//idP, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	return c.JSON(http.StatusNotFound, errors.New("your requested Item is not found"))
	//}
	//
	//id := int64(idP)
	//ctx := c.Request().Context()
	//
	//art, err := a.PLogic.GetByID(ctx, id)
	//if err != nil {
	//	return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	//}
	//
	//return c.JSON(http.StatusOK, art)
	return nil
}

func (a *PersonHandler) Update(c echo.Context) error {
	//idP, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	//return c.JSON(http.StatusNotFound, app.ErrNotFound.Error())
	//}
	//
	//id := int64(idP)
	//ctx := c.Request().Context()
	//
	//err = a.PLogic.Delete(ctx, id)
	//if err != nil {
	//	return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	//}
	//
	//return c.NoContent(http.StatusNoContent)
	return nil
}

func (a *PersonHandler) Delete(c echo.Context) error {
	//idP, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	//return c.JSON(http.StatusNotFound, app.ErrNotFound.Error())
	//}
	//
	//id := int64(idP)
	//ctx := c.Request().Context()
	//
	//err = a.PLogic.Delete(ctx, id)
	//if err != nil {
	//	return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	//}
	//
	//return c.NoContent(http.StatusNoContent)
	return nil
}

func getStatusCode(err error) int {
	//if err == nil {
	return http.StatusOK
	//}

	logrus.Error(err)
	switch err {
	case errors.New("internal Server Error"):
		return http.StatusInternalServerError
	case errors.New("your requested Item is not found"):
		return http.StatusNotFound
	case errors.New("your Item already exist"):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
