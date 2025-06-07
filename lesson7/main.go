package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Species string 

}

type Bird struct {
	Name string 
	Species string 
}

type Animal interface{
	MakeSound() string 
}

func printAnimalSound(a Animal){

	fmt.Println(a.MakeSound())

}

func main() {

	bird := Bird{Name: "Kesha", Species: "Parrot"}
	dog := Dog{Name: "Boris", Species: "Retriver"}

	printAnimalSound(bird)
	printAnimalSound(dog)
}
