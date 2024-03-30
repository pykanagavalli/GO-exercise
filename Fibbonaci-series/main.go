package main

import "fmt"

func main() {
	mychan := make(chan int)
	go Fibbonaci(mychan, 5)
	for val := range mychan {
		fmt.Println(val)
	}
}

func Fibbonaci(mychan chan int, limit int) {
	sum := 0
	a := 1
	b := 0

	for i := 0; i <= limit; i++ {
		sum = a + b
		a, b = b, sum

		mychan <- sum
	}
	close(mychan)

}
