package main

import (
	"fmt"
)

func main() {
	array := [6]int{5, 6, 8, 9, 3, 1}

	max := array[0]
	min := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]

		}
		if array[i] < min {
			min = array[i]
		}

	}

	fmt.Println("maximum", max)
	fmt.Println("minimum", min)
}
