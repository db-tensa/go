package main

import (
	"fmt"               // you know why
	"golang.org/x/term" // importing it for enter in a raw mode in terminal
	"os"                // to get an file access
	"strings"           // for castring from one data type to another
)

func main() {

	// -- my program requires an additional file, a "crutch file", so, I make my programm accept args
	// the second arg must be a file name which alocated in your directory
	if len(os.Args) < 2 {
		fmt.Println("Not enough args")
		fmt.Println("your input must be: go run main.go filename")
		os.Exit(1)
	}

	filename := os.Args[1] // getting a filename, in future this variable will be used as a more simple way to access a file

	// reading a file , this is a high-levle reading acess, ofcourse, I could made it with syscall, but there is no reason to use it
	writed_text, err := os.ReadFile(filename) // reading a file
	if err != nil {                           // checking if it's okay
		fmt.Println("Error getting a file. Check your directory permissions, and check file itself")
		os.Exit(1)
	}

	// converting a file output from byte format to string format.
	// And i made it more clener, Because of new implementaion of the loop in the "buffer". I don't need any more loop which adds words in slice
	original_text := string(writed_text)

	//This string is a kind of support, which allows me to use input with a short reaction of the program to the text entered by the user,
	//which, by the way, saves me from some small problems. But on the other hand, it prevents me from entering characters,
	//why - I could not figure out, I tested it in other programs, there it worked fine, here - a bug.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	// it preventing a transormation  your terminal in to the garbage
	// Btw I forget to mention it in my first homework
	// Defer like a list, it will be executed roughly speaking almost before closing the program
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// we will use a buffer, as our container for words
	buffer := []byte{}

	// we write an \r to format the string
	fmt.Println("\r", "Hello ! This a program, it's verion is 0.1 i will make it better in the future")
	fmt.Println("\r", "------------------------")

	fmt.Println("\r", "You must execute a bash script for proper work, also, if you wanna shut the programm, press the ESC")
	fmt.Println("\r", "------------------------")

	fmt.Println("\r", "GOOD LUCK")
	fmt.Println("\r", "------------------------")

	// printing original text before any user input
	fmt.Print("\033[2J\033[H")
	fmt.Println(original_text)

	for {
		// this variable will contain an one symbol which user writed
		buf := make([]byte, 1)
		_, err := os.Stdin.Read(buf) // checking if our buffer is ok !
		if err != nil {
			panic(err)
		}

		// points to the beginning of the buffer for proper operation
		b := buf[0]

		// for program shutting (27 == escape)
		if b == 27 {
			fmt.Print("\033[2J\033[H")
			fmt.Println("Escape pressed, shut down the program...")
			defer fmt.Print("\033[2J\033[H")
			os.Exit(0)
		}

		// determinate the backspace
		if b == 127 || b == 8 {
			if len(buffer) > 0 {
				// basically, I found the only way - after each backspace that user pressed,
				//we creating a new buffer, which size is smaller than previous buffer size
				buffer = buffer[:len(buffer)-1]
			}
		} else if b >= 32 && b <= 126 {
			// checking the letter (ASCII LETTER ! ) that user enter
			buffer = append(buffer, b) // pushing the word in to the massive
		}

		input := string(buffer) // again, transform it form byte format to string format

		// clearing a window
		fmt.Print("\033[2J\033[H")
		// new output
		fmt.Println("\r\x1b[34m >> ", input, " << \x1b[0m")

		// comfortable ? 
		highlighted_text := original_text

		if len(input) > 0 { // i know i know,  that i should check it earlier, but still good
			words := strings.Fields(original_text)
			for _, word := range words {
				if strings.HasPrefix(word, input) {
					// HERE IS THE COLORFUL TEXT !
					// and yeah, I make it more simplier and more attractive. Of course, I could do it through a library with color styles, but ascii code are just simplier.
					highlighted := "\033[31m" + word + "\033[0m"
					highlighted_text = strings.ReplaceAll(highlighted_text, word, highlighted)
				}
			}
		}



		// printing the beauty text 💅💅💅💅💅💅💅💅
		fmt.Println("\r\x1b[32m\x1b[1m<<<<<<<< Your text >>>>>>>>\x1b[0m")
		fmt.Println("\r", "\n", highlighted_text)
		fmt.Println("\r\x1b[32m\x1b[1m<<<<<<<< End of text >>>>>>>>\x1b[0m")
		//of course I had realization where all words were adding in the slice EXPECT the last one, so here is the "uga buga" realization
		//Basically it's just breaking the whole text in words 
		words := strings.Fields(original_text)
		fmt.Println("\r","Amount of words",len(words))
	}
}
