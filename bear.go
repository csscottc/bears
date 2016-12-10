package main

import "fmt"

//bear structure - bears can have a hat !
//uses composition on animalStruct
type bearStruct struct {
	hasHat bool
	animalStruct
}

//this is a constructor function for a bear
func newBearStruct(name string, age int) *bearStruct {
	result := bearStruct{} //create an instance of a bear
	result.name = name
	result.age = age
	result.hasHat = true
	return &result //return the memory address of the bear we created !
}

//this adds the printHatStatus method to the bear struct
func (bear *bearStruct) printHatStatus() {
	fmt.Println("Does the bear have a hat? ", bear.hasHat)
}
