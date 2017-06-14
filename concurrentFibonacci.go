package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int)
	go goFib(n, c)
	return c
}

func goFib(n int, c chan int) {
	if n<2 {
		c <- 1
	} else {
		x := <-fib(n-1)
		y := <-fib(n-2)
		c <- x+y
	}
	close(c)
}

func main() {
	fmt.Println("fibonacci: ", <-fib(5))
}