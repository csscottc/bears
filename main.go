package main

import "fmt"

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
