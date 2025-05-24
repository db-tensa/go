package main

import (
	"fmt"
)

func main() {
	fmt.Println("Yo")
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	var fr [5]string = [5]string{"orange", "banana", "peach", "apple", "grapes"}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(fr[i])
	}

}
