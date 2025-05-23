package main

import "fmt"


	func inctement() func() int {
	count := 0 
	return func() int{
		count++ 
		return count
	}
}
func main() {

	inc := inctement()
	fmt.Println(inc())

}


