package main

import "fmt"

func main() {
	// crearting map and check if nil or not
	var map1 map[int]int
	if map1 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// Now we create new map and give some values

	map2 := map[int]string{
		1: "ganga",
		2: "preetha",
		3: "sowmi",
	}
	fmt.Println(map2)
}
