package handler

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	validate = validator.New()
)

type APIHander struct {
}

func NewApiHandler() *APIHander {
	return &APIHander{}
}

type AddInput struct {
	num  []int  `query:"num" validate:"required,gt=1"`
	hoge string `query:"hoge" validate:"required"`
}

func (i AddInput) Validate() error {
	return validate.Struct(i)
}

func (h *APIHander) Add(c echo.Context) error {
	var param AddInput
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := param.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println(param)
	var sum int
	for i := range param.num {
		log.Println(i)
		sum += i
	}
	log.Println(c.QueryParam("hoge"))
	return c.JSON(http.StatusOK, sum)
}
