package main

import (
	"fmt"
	"sync"
)

func say(n int) {
	fmt.Println(fmt.Sprintf("hello # %d", n))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			say(i)
		}(i)
	}
	wg.Wait()

	foo()
label:
	fmt.Println("goto here")
}

func foo() {
label:
	fmt.Println("foo")
	goto label
}
