package main

import "fmt"

func main(){

	m := make(map[string]struct{
		X int
		Y string
		Z rune
	})


	m[X]=10

	fmt.Println(m["key"])

	delete(m, "key")

	v, exists := m["key"]

	if exists {
		fmt.Println(v)
	} else{
		fmt.Println("deleted")
	}
}
