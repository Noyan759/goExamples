package main

import "fmt"

func main() {

	fmt.Println("hello world")

	a, b := 1, 2

	a, b = 5, a+b

	var x uint64 = 1<<64-1

	fmt.Println(a, b)
	fmt.Println(x)

	y:=(1.414*1.414+2)/(2*1.414)
	fmt.Println(y)

}
