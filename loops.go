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

	sliice :=[]string{"Alex", "Nikita", "Andrey"}

	// i := 0
	// for  i < 3 {
	// 	a := i+3
	// 	fmt.Println("Iteration ", i, i, i, a)
	// }

	for i := 0; i < len(sliice); i++{

		fmt.Println(sliice[i], i )
	}
	outerLoop:
	for i := 0; i < 4; i++{
		for j:= 0; j < 4; j++{
			fmt.Println("Error")
			break outerLoop
		}

	}
	for i := 0; i < 10; i++{
		if i % 2 == 0{
			continue
		}else{
			fmt.Println("Even", i)
		}
	}


	for i := 2; i < 10; i++ {
		
		if 21%i == 0  || i / 9 == 1{
			fmt.Println(i)
		}
	}


	fmt.Println("New ")
	for i := 0; i < 10; i++ {
    if i > 4 && i < 7 {
        continue
    }
    fmt.Println("число:", i)
}

}

