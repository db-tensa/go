package main

import (
	"fmt"
)

type mover interface{
	Move() string
}
type Robot struct{
	Model string 
}

func (r Robot ) Speak() string{
	return "BEep beep"
}

func (r Robot) mover() string{
	return "how wheel"
}

func main(){
	robot := Robot{Model: "R02"}
}

