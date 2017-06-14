package main

import (
	"time"
	"fmt"
)

func main() {

	tick := time.Tick(100*time.Millisecond)
	boom := time.After(500*time.Millisecond)

	defer fmt.Println("blahblahblah")

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			fmt.Println(len(tick))
			goto someLabel
		default:
			time.Sleep(50*time.Millisecond)
			fmt.Println(".")
		}
	}

	someLabel:
	fmt.Println("goto here")
}