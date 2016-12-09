package main

import (
	"fmt"
)

func main() {
	bear := newBearStruct("Yogi", 42)
	koala := newKoalaBearStruct("Gary", 21, true, true)
	fmt.Println("Welcome to the bears application")
	fmt.Println("Our bear:", *bear)
	fmt.Println("Our koala:", *koala)
	fmt.Println("--------------------------------")
	bear.printDetails()
	bear.printHatStatus()
	fmt.Println("--------------------------------")
	koala.printDetails()
	koala.printHatStatus()
	koala.printTieStatus()
}

//koala bear structure - koala bears can have a hat and a tie !
type koalaBearStruct struct {
	hasTie bool
	bearStruct
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

//yes, I am fully aware that a koala is a marsupial and not a bear.
func newKoalaBearStruct(name string, age int, tie bool, hat bool) *koalaBearStruct {
	koala := koalaBearStruct{}
	koala.age = age
	koala.name = name
	koala.hasHat = hat
	koala.hasTie = tie
	return &koala
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

//print out the tie status
func (koala *koalaBearStruct) printTieStatus() {
	fmt.Println("Does the koala have a tie? ", koala.hasTie)
}
