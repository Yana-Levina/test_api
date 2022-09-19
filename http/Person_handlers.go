package http

import (
	"github.com/labstack/echo/v4"
	_ "net/http"
	_ "strconv"
	"test/app"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
//type ArticleHandler struct {
//	AUsecase domain.ArticleUsecase
//}

func NewPersonHandler(e *echo.Echo, us app.PersonUsecase) {
	handler := &PersonUsecase{
		PUsecase: us,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}

func createPerson() {

}

func GetByID() {

}

func Update() {

}

func Delete(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	//delete(users, id)
	//return c.NoContent(http.StatusNoContent)
	return nil
}
