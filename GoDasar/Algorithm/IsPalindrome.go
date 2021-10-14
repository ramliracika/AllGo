package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool  {
	var res,temp string

	temp = strings.ToLower(str)
	for i:= len(str)-1; i >= 0; i--{
		res += strings.ToLower(string(str[i]))
	}

	if res != temp {
		return false
	}

	return true


}

func main () {
	fmt.Println(IsPalindrome("AbBa"))
}
