package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type RelationInfo struct {
	Influenced_by []string `json:"influenced-by"`
	Influences []string `json:"influences"`
}

type LanguagesInfo struct {
	Language string `json:"language"`
	Appeared int `json:"appeared"`
	Created []string `json:"created"`
	Functional bool `json:"functional"`
	Object_oriented bool `json: "object-oriented"`
	Relation RelationInfo `json: "relation"`
}

func get_language_mock_data() LanguagesInfo {
	return LanguagesInfo{
		Language: "Go",
		Appeared: 2009,
		Functional : true,
		Object_oriented: false,
		Created: []string{"Robert Griesmer", "Rob Pike", "Ken Thompson"},
		Relation: 	RelationInfo{
				Influenced_by: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"}, 
				Influences: []string{"C++", "Objective-C", "C#", "Java", "Javascript", "PHP", "Go"},
			},
	}
}

func LanguageController(c echo.Context) error {
	mock_data := get_language_mock_data()
	return c.JSON(http.StatusOK, mock_data)
}