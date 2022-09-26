package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"test/app"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
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

func (p *PersonHandler) Create(c echo.Context) (err error) {
	var person app.Person

	if err := c.Bind(&person); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&person); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = p.PUsecase.Create(ctx, &person)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	log.Printf("res %+v", person)
	return c.JSON(http.StatusCreated, person)
}

func (p *PersonHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	listAr, err := p.PUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

func (p *PersonHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Your requested person is not found"))
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := p.PUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func (p *PersonHandler) Update(c echo.Context) (err error) {
	var person app.Person

	if err := c.Bind(&person); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&person); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Your requested person is not found"))
	}

	ctx := c.Request().Context()
	err = p.PUsecase.Update(ctx, &person, int64(id))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, person)
}

func (p *PersonHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Your requested person is not found"))
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = p.PUsecase.Delete(ctx, id)
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
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case errors.New("Internal Server Error"):
		return http.StatusInternalServerError
	case errors.New("Your requested person is not found"):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
