package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
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

func NewPersonHandler(e *echo.Echo, pu app.PersonUsecase) {
	handler := &PersonHandler{
		PUsecase: pu,
	}
	e.POST("/persons", handler.Create)
	e.GET("/persons", handler.GetAll)
	e.GET("/persons/:id", handler.GetByID)
	e.PUT("/persons/:id", handler.Update)
	e.DELETE("/persons/:id", handler.Delete)
}

func (a *PersonHandler) Create(c echo.Context) (err error) {
	var person app.Person

	if err := c.Bind(&person); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&person); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.PUsecase.Create(ctx, &person)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	//log.Printf("req data %+v", person)
	return c.JSON(http.StatusCreated, person)
}

func (a *PersonHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

func (a *PersonHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("your requested Item is not found"))
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.PUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
	//return nil
}

func (a *PersonHandler) Update(c echo.Context) (err error) {
	var person app.Person
	err = c.Bind(&person)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&person); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.PUsecase.Update(ctx, &person)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, person)
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

func isRequestValid(m *app.Person) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
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
