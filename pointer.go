package main

import "fmt"

func main() {

	foo1, foo2 := 1, 2

	fmt.Println(foo1, foo2)

	swap(&foo1, &foo2)

	fmt.Println(foo1, foo2)

}

func swap(foo1, foo2 *int){
	temp:=*foo1
	*foo1=*foo2
	*foo2=temp
}