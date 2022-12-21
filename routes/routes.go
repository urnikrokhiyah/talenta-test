package routes

import (
	"github.com/labstack/echo/v4"
	"talenta-test/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", controllers.TestController)
	e.GET("/palindrome", controllers.PalindromeController)
	return e
}