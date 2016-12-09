package main

import (
	"fmt"
)

func main() {
	bear := newBearStruct("Yogi", 42)
	fmt.Println("Welcome to the bears application")
	fmt.Println("Our bear:", *bear)
	fmt.Println("--------------------------------")
	bear.printDetails()
	bear.printHatStatus()
}

//bear structure - bears can have a hat !
//uses composition on animalStruct
type bearStruct struct {
	hasHat bool
	animalStruct
}

//base struct for animals - animals have a name and age.
type animalStruct struct {
	name string
	age  int
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

//this prints some details about an animal
func (animal *animalStruct) printDetails() {
	fmt.Println("Name:", animal.name)
	fmt.Println("Age:", animal.age)
}
