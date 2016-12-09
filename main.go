package main

import (
	"fmt"
)

func main() {
	bear := newBearStruct("Yogi")
	fmt.Println("Welcome to the bears application")
	fmt.Println(bear.name)
}

//This is the definition of the bear structure
type bearStruct struct {
	name string
}

//This is a constructor function for a bear
func newBearStruct(name string) *bearStruct {
	result := bearStruct{} //Create an instance
	result.name = name
	return &result //de-reference
}
