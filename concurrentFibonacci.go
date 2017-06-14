package main

import (
	"fmt"
)

func goconcfib(n int) chan int {
	result := make(chan int)
	go concfib(n, result)
	return result
}

func concfib(n int, output chan int) {
	if n < 2 {
	output <- 1
	} else {
	a, b := goconcfib(n-1), goconcfib(n-2)
	output <- (<-a) + (<-b)
	}
	close(output)
}

func main() {
	number:= 50
	fmt.Printf("fib(%d) = %d\n", number, <-goconcfib(number))
}