package main

import "fmt"

//base struct for animals - animals have a name and age.
type animalStruct struct {
	name string
	age  int
}

//this prints some details about an animal
func (animal *animalStruct) printDetails() {
	fmt.Println("Name:", animal.name)
	fmt.Println("Age:", animal.age)
}
