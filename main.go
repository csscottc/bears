package main

import (
	"fmt"
)

func main() {
	bear := newBearStruct("Yogi")
	fmt.Println("Welcome to the bears application")
	bear.printName()
	fmt.Println(*bear)
}

//This is the definition of the bear structure
type bearStruct struct {
	name string
	age  int
}

//This is a constructor function for a bear
func newBearStruct(name string) *bearStruct {
	result := bearStruct{} //Create an instance
	result.name = name
	result.age = 42
	return &result //de-reference (Return the bear, not the memory address)
}

func (bear *bearStruct) printName() {
	fmt.Println(bear.name)
}
