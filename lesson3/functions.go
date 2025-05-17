package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(divide(10, 2))

	_, err := divide(10, 2)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("ok")

	}

}

func getData() (int, string, bool) {
	return 42, "hello", true
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Zero divide")
	}
	return a / b, nil
}
