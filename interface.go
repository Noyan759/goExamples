package main

import "fmt"

type Dog struct {}

func (d Dog) Speak()string {
	return "woof woof"
}

type Cat struct {}

func (c *Cat) Speak()string {
	return "meow meow"
}

func (c *Cat) RollOnTheFloor() {
	fmt.Println("rolling...")
}

type Animal interface {
	Speak() string
}

func main(){
	animals := []Animal{
		&Dog{},
		&Cat{},
	}

	for _,v := range animals{
		fmt.Println(v.Speak())
	}

	c := animals[1]

	cat, ok := c.(*Cat)
	if ok{
		cat.RollOnTheFloor()
		fmt.Println("ok")
	} else{
		fmt.Println("not ok")
	}
}