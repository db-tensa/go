package main // special package for entry point

import (
	"fmt"
	"unsafe"
)

func main() { // entry point

	const PI = 3.14
	const smth = "smth"

	var _flooat32 float32 = 32.31

	var _a int64 = 8
	var _b int64 = 1

	_sum_of := _a + _b

	var _a_dis int64 = 8
	var _b_dis int64 = 1

	_dis_of := _a_dis - _b_dis

	var _name1 string             // first method
	_new := "I really don't know" // second method
	var _name2 string = "smth"    // third method
	_name1 = "idk"

	// if doing variable  := data  this way
	//then you cannot set a type of variable even a const

	floatQuo := float64(_a) / float64(_b)
	if true {
		value := 20
		fmt.Println("\n This is the true/false output variable -> ", value)
	}
	// var (
	// 	firstName string = "Alex"
	// 	lastName string = "Skyba"
	// 	year int = 1000;
	// )
	fmt.Println("\n This is name 1 -> ", _name1, " \n This is the name 2 -> ", _new, " \n This is the name 3 -> ", _name2)

	fmt.Println("\n Size of int is ", unsafe.Sizeof(int(0)))
	fmt.Println("\n This is PI (const) -> ", PI, "\n This is smth -> (const) ", smth)
	fmt.Println("\n This is float32 -> ", _flooat32)
	fmt.Println("\n Size of string is ", unsafe.Sizeof(_name2), "byte")

	fmt.Println("This is a sum of a and b ->", _sum_of)
	fmt.Println("This is a dis of a and b ->", _dis_of)
	fmt.Println("This is a type casting -> ", floatQuo)

}
