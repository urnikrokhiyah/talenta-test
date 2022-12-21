package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
)

type RelationInfo struct {
	Influenced_by []string `json:"influenced-by"`
	Influences []string `json:"influences"`
}

type LanguageInfo struct {
	ID int `json:"id"`
	Language string `json:"language"`
	Appeared int `json:"appeared"`
	Created []string `json:"created"`
	Functional bool `json:"functional"`
	Object_oriented bool `json: "object-oriented"`
	Relation RelationInfo `json: "relation"`
}

var languages_info []LanguageInfo = []LanguageInfo{
	LanguageInfo{
		ID: 1,
		Language: "Go",
		Appeared: 2009,
		Functional : true,
		Object_oriented: false,
		Created: []string{"Robert Griesmer", "Rob Pike", "Ken Thompson"},
		Relation: 	RelationInfo{
			Influenced_by: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"}, 
			Influences: []string{"C++", "Objective-C", "C#", "Java", "Javascript", "PHP", "Go"},
		},
	},
}

func GetLanguagesListController(c echo.Context) error {
	return c.JSON(http.StatusOK, languages_info)
}

func GetLanguageController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mapping_language := map[int]LanguageInfo{}
	for i :=0 ; i < len(languages_info); i++ {
		mapping_language[languages_info[i].ID] = languages_info[i]
	}
	
	if _, value := mapping_language[id]; value {
		return c.JSON(http.StatusOK, mapping_language[id])
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "not found"})
}

func DeleteLanguageController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mapping_language := map[int]LanguageInfo{}
	for i :=0 ; i < len(languages_info); i++ {
		mapping_language[languages_info[i].ID] = languages_info[i]
	}

	if _, value := mapping_language[id]; value {
		delete(mapping_language, id)
		var new_language_list []LanguageInfo
		for _, value := range mapping_language {
			new_language_list = append(new_language_list, value)
		}

		languages_info = new_language_list
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "success"})

	}
	
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "not found"})
}

func CreateLanguageController(c echo.Context) error {
	var new_language LanguageInfo
	err := c.Bind(&new_language)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "failed"})
	}

	new_language.ID = len(languages_info) + 1
	languages_info = append(languages_info, new_language)
	return c.JSON(http.StatusOK, languages_info)
}

func UpdateLanguageController(c echo.Context) error {
	var new_language LanguageInfo
	err := c.Bind(&new_language)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "failed"})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	mapping_language := map[int]LanguageInfo{}
	for i :=0 ; i < len(languages_info); i++ {
		mapping_language[languages_info[i].ID] = languages_info[i]
	}

	if _, value := mapping_language[id]; value {
		if new_language.ID == 0 {
			new_language.ID = id
		}

		mapping_language[id] = new_language
		var new_language_list []LanguageInfo
		for _, value := range mapping_language {
			new_language_list = append(new_language_list, value)
		}

		languages_info = new_language_list
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "success"})

	}
	
	return c.JSON(http.StatusBadRequest, map[string]string{"message": "not found"})
}