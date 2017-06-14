package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//go say("world")
	//say("hello")

	//for i:=0; i<100; i++ {
	//	go func (i int){
	//		fmt.Println(i)
	//	}(i)
	//}

	c := make(chan int)

	go func () {
		time.Sleep(time.Second)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	fmt.Println("RUN")

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Wait")
}