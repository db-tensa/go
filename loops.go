package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("CONDITIONS")
	fmt.Println("--------------------------------------------------------------")
	currentHout := time.Now().Hour()
	fmt.Println(currentHout)
	if currentHout >= 6 && currentHout < 22 {
		fmt.Println("Day")
	} else {
		fmt.Println("Night")
	}

	fmt.Println("LOOOPS")
	fmt.Println("--------------------------------------------------------------")


	for i := 0; i < 3; i++ {
		a := i+3
		fmt.Println("Iteration error", i, i, i, a)
	}
}
