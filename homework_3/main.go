package main

import (
	"fmt"               // you know why
	"golang.org/x/term" // importing it for enter in a raw mode in terminal
	"os"                // to get an file access
	"strings"           // for castring from one data type to another
	"unicode"           // for checking a symbols
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
	// We need  it because of next loop iteration.
	user_text := string(writed_text)

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

	// we converting the text from file in to the slice,  because
	// 1. Using  a slice, counting words - the simple task
	// 2. With slices I can eaaaaaaaaaaasily access to needed word without any "troubels"
	users_words_in_slice_mode := []string{}

	var word_upon_a_space string // if in future, the iterator will face to face a space a.k.a a " ", then the word he readed will be allocated in slice

	

	// iterating on a text
	for _, iter := range user_text {
		// checking the size, if 0 then return 1 
		if len(user_text) == 0{
			fmt.Println("Variable is empty shutting the program")
			os.Exit(1)
		}
				// checking if our symbol is a letter

		if unicode.IsLetter(iter) {
			// adding a word into a "sub-variable"
			word_upon_a_space += string(iter)
		} else if iter == ' ' {

			// meet a space ? Add a word in slice
			users_words_in_slice_mode = append(users_words_in_slice_mode, word_upon_a_space)
			// cleaning it, if we dont do it, then the word we added before, will also be in the slice, here's an example "fire fire hello fire no fire what"
			word_upon_a_space = ""
		}

	}


	// all words
	all_words := len(users_words_in_slice_mode)

	
	// ahh, just for beauty : )
	var horiz int
	for i := 0; i < len(users_words_in_slice_mode); i++ {
		horiz += i * 2 
	}



	// we will use a buffer, as our container for words
	buffer := []byte{}

	// we write an \r to format the string
	fmt.Println("\r", "Hello ! This a program, it's verion is 0.1 i will make it better in the future")
	fmt.Println("\r", "------------------------")

	fmt.Println("\r", "You must execute a bash script for proper work, also, if you wanna shut the programm, press the ESC")
	fmt.Println("\r", "------------------------")

	fmt.Println("\r", "GOOD LUCK")
	fmt.Println("\r", "------------------------")

	for {
		// this variable will contain an one symbol which user writed
		buf := make([]byte, 1)
		_, err := os.Stdin.Read(buf) // checking if our buffer is ok !
		if err != nil {
			panic(err)
		}

		// points to the beginning of the buffer for proper operation
		b := buf[0]

		// determinate the backspace
		if b == 127 || b == 8 {
			if len(buffer) > 0 {
				// basically, I found the only way - after each backspace that user pressed,
				//we creating a new buffer, which size is smaller than previous buffer size
				buffer = buffer[:len(buffer)-1]
			}
			// checking the letter (ASCII LETTER ! ) that user enter
		} else if b >= 32 && b <= 126 {
			buffer = append(buffer, b) // pushing the word in to the massive
		}

		// for program shutting (27 == escape)
		if b == 27 {
			fmt.Print("\033[2J\033[H")
			fmt.Println("Escape pressed, shut down the program...")
			defer fmt.Print("\033[2J\033[H")
			os.Exit(0)
		}

		// clearing a window
		fmt.Print("\033[2J\033[H")

		input := string(buffer)                   // again, transform it form byte format to string format
		fmt.Println("Enter something >> ", input) // there is a buffer !


		// iterati on user input
		for _, word := range users_words_in_slice_mode {
			if strings.HasPrefix(word, input) { // find a match ? Print it !
				fmt.Println("\r", word) 
				fmt.Println("\r", "------------------------")
			}
		}

		// for beauty  ^_^. What I just typed
		for i := 0; i < horiz; i++ {
			fmt.Print("\r", "-")
		}

		// So the user can see the word or the letter he inputed, and also the text !
		fmt.Println("\r", user_text)
		// printing an amount of words that user input contains 
		fmt.Println("\r", "Your text has ", all_words, "words")

	}
}
