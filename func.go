package main

import "fmt"


func palindrome(word string) bool {
	var length_word = len(word)
	for i:=0; i < length_word; i++{
		if word[i] != word[length_word-1-i]{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("kasur rusak"))
}