package main

import (
	"fmt"
	"time"
	"unicode"
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


	outerLoop:
	for i := 0; i < 4; i++{
		for j:= 0; j < 4; j++{
			fmt.Println("Error")
			break outerLoop
		}

	}


	for i := 2; i < 10; i++ {

		if 21%i == 0  || i / 9 == 1{
			fmt.Println(i)
		}
	}


	var ch byte = 'A'
	fmt.Println(string(ch))

var name = "dsds"
	for _, r := range name{
		if !unicode.IsLetter(r) && !unicode.IsSpace(r){
			fmt.Println("bad")
		}
	}
	text := "Hello dsds"
	runes := []rune(text)
	fmt.Println(runes)

	s := "ID彼氏彼女の事情"

	fmt.Println(len(s))

	for i:= 0; i < len(s); i ++{
		fmt.Println( "%x" ,s[i])
	}
}

