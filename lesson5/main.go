package main

import (
	"fmt"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	slice1 := arr[1:4]
	slice2 := arr[0:3]

	fmt.Println(slice1)
	fmt.Println(slice2)

}
