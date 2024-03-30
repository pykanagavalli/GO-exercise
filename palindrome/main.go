package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "121"
	//if the string are palindrome are not
	if reverse(strings.Trim(str, "")) == strings.Trim(str, "") {
		fmt.Println(str + "\nIt is a palindrome")
	} else {
		fmt.Println(str + "\nIt is not a palindrome")
	}
}

func reverse(str string) (result string) {
	for _, val := range str {
		result = string(val) + result
	}
	return result
}
