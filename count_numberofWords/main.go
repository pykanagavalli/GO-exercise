package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This world is very cruiel"
	res := strings.Split(text, "")
	fmt.Println(len(res))
}
