package main // special package for entry point

import(
	"fmt"
	"unsafe"
)

func main() { // entry point

	
	var _name1 string;

	 _name1 = "idk"
	// var (
	// 	firstName string = "Alex"
	// 	lastName string = "Skyba"
	// 	year int = 1000;
	// )
	fmt.Println("Hello, world.") // output
	fmt.Println(_name1);
	fmt.Println("Size of int is ", unsafe.Sizeof(int(0)))

}
