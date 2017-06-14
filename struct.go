package main

import "fmt"

type Foo struct {
	x int
	y string
}

func (f *Foo) inc (){
	f.x++
}

func main() {
	f:=Foo{15, "oooof"}

	fmt.Println(f)

	f.x=1
	f.y="foooo"

	fmt.Println(f)

	f.inc()

	fmt.Println(f)

}