package main

import "fmt"

func sum(s []int, c chan int) {
	// implement a method that sums all elements of `s` slice and
	// sends the result to `c` channel
	lSum := s[0] + s[1] +s[2]
	c <- lSum
}

func main() {
	s := []int{1, 5, -2, 7, 2, 8, -9, 4, 0}
	c := make(chan int)

	//lSlice := s[0:3]
	//go sum(lSlice, c)
	//
	//n := <- c
	//
	//fmt.Println(n)

	for i:=0; i<len(s); i=i+3{
		//fmt.Println(i)
		lSlice := s[i:i+3]
		go sum(lSlice, c)
		//fmt.Println(len(s), lSlice, c)

	}
	//close(c)
	totalSum := 0

	x := <- c
	y := <- c
	z := <- c
	totalSum = x+y+z

	//for v := range c {
	//	fmt.Println("local sum: %d", v)
	//	totalSum = totalSum + v
	//	fmt.Println("hello", totalSum)
	//}

	fmt.Println("Total sum is %d", totalSum)

	// sum `s` slice by taking three elements each time
	// print sum of each time and the total result
}