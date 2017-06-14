package main

import "fmt"

func main(){

	slice := [] int{}

	fmt.Println(slice)

	s := slice[:]

	fmt.Println(s)

	for i:=0; i<10; i++ {
		slice = append(slice, i)
		printSlice(slice)
	}
}

func printSlice(slice []int){
	fmt.Println("len=%d cap=%d array=%v", len(slice), cap(slice), slice)
}