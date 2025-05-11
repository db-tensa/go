package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func main() {

	terminal := term.IsTerminal(int(os.Stdin.Fd()))
	if !terminal {
		os.Exit(1)
		fmt.Print("Something wrong")
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("Error when trying to save old terminal state")
		os.Exit(1)

	}
	// adding a term.Restore to the end of the "defer" list to execute it when programm end
	// to keep the terminal in a good "shape"
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	syscall.SetNonblock(int(os.Stdin.Fd()), true)

	// user input text for finding a word / initial letter of the word
	var user_text string
	// creating a buffer for

	buffer := [148]byte{}
	for  {

		buf := make([]byte, 148)
		_, err = os.Stdin.Read(buf)
	
	}
	

}
