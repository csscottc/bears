package main

import "fmt"

//koala bear structure - koala bears can have a hat and a tie !
type koalaBearStruct struct {
	hasTie bool
	bearStruct
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

//print out the tie status
func (koala *koalaBearStruct) printTieStatus() {
	fmt.Println("Does the koala have a tie? ", koala.hasTie)
}
