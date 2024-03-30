// create factoriral with using channel and go routine

package main

import "fmt"

func main() {
	mychan := make(chan int)

	go Factorial(mychan, 10)
	for val := range mychan {
		fmt.Println(val)
	}
}

func Factorial(mychan chan int, limit int) {
	res := 1
	for i := 1; i <= limit; i++ {
		res *= i
	}
	mychan <- res
	close(mychan)
}

// out put
// 3628800
// 10*9*8*7*6*5*4*3*2*1 = 3628800
