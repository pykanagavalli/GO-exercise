package main

import (
	"fmt"
	"os"
)

func main() {
	res := os.Args
	fmt.Println(res[2])
}
