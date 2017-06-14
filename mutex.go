package main

import (
	"fmt"
	"sync"
)

type Integer struct {
	value int
	m sync.Mutex
}

func (i *Integer) Inc() {
	i.m.Lock()
	i.value++
	i.m.Unlock()
}

func (i *Integer) Val() int {
	return i.value
}

func main(){
	n:= &Integer{}

	var wg sync.WaitGroup
	wg.Add(10000)

	for i:=0; i<1000; i++ {
		go func(){
			defer wg.Done()
			n.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(n.Val())
}
