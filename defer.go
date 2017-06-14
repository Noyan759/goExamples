package main

import "fmt"

func main() {


	defer fmt.Println("bye world")

	defer func(){
		err:=recover()
		fmt.Println(err)
	}()

	fmt.Println(("hello world"))

	panic("boo")

	defer fmt.Println("bye bye world")

	fmt.Println(("hello hello world"))

	defer fmt.Println("bye bye bye world")

}
