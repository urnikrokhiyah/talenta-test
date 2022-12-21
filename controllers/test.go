package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func TestController(c echo.Context) error {
	result := map[string]string{"message": "Hello Go developers!"}
	return c.JSON(http.StatusOK, result)
}