package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
	Group string 
	Grades []int
	Average float64 
}
type Rectangle struct{
	Width float64
	Height float64 
}
func (s Student) AddGreade (grade float64){
	s.Grades = append(s.Grades, grade)

}

func (r Rectangle )Area() float64{
	return r.Width * r.Height 
}

func main() {

	
}
