package main

import (
	"fmt"
)

func main() {
	bear := newBearStruct("Yogi")
	fmt.Println("Welcome to the bears application")
	fmt.Println(*bear)
	fmt.Println("--------------------------------")
	bear.printName()
	bear.printAge()
}

//bear structure
type bearStruct struct {
	hasHat bool
	animalStruct
}

//base struct for animals
type animalStruct struct {
	name string
	age  int
}

//This is a constructor function for a bear
func newBearStruct(name string) *bearStruct {
	result := bearStruct{} //Create an instance
	result.name = name
	result.age = 42
	result.hasHat = true
	return &result //de-reference (Return the bear, not the memory address)
}

func (bear *bearStruct) printName() {
	fmt.Println("Does the bear hava hat:", bear.hasHat)
}

func (animal *animalStruct) printAge() {
	fmt.Println("Name:", animal.name)
	fmt.Println("Age:", animal.age)
}
