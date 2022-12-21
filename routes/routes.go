package routes

import (
	"github.com/labstack/echo/v4"
	"talenta-test/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", controllers.TestController)
	e.GET("/palindrome", controllers.PalindromeController)
	e.GET("/languages", controllers.GetLanguagesListController)
	e.POST("/language", controllers.CreateLanguageController)
	e.GET("/language/:id", controllers.GetLanguageController)
	e.DELETE("/language/:id", controllers.DeleteLanguageController)
	e.PATCH("/language/:id", controllers.UpdateLanguageController)
	return e
}