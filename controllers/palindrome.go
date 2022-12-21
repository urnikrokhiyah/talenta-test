package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func palindrome_check(word string) bool {
	length_word := len(word)
	for i := 0; i < length_word; i++ {
		if word[i] != word[length_word-1-i] {
			return false
		}
	}

	return true
}

func get_palindrome_response(is_palindrome bool) (int, map[string]string) {
	if is_palindrome == true {
		return http.StatusOK, map[string]string{"message": "Palindrome"}
	}

	return http.StatusBadRequest, map[string]string{"message": "Not Palindrome"}
}

func PalindromeController(c echo.Context) error {
	var palindrome_word string = c.QueryParam("text")
	if palindrome_word == ""{
		http_response, result := get_palindrome_response(false)
		return c.JSON(http_response, result)
	}

	is_palindrome := palindrome_check(palindrome_word)
	http_response, result := get_palindrome_response(is_palindrome)
	return c.JSON(http_response, result)
}

