// I wanna say that I am not proud of this project, the FIRST one was and still better then this, maybe because I do it at the night time, or maybe because I don't "invent" a bycycle this time.
// But still, I can say that code is looking a little bit better then in first one, but realization is kinda boring
package main

import (
	"bufio"   // for grabbing user input like a pro
	"fmt"     // you know, for printing stuff
	"os"      // for stdin, 'cause we need user input now
	"strings" // for messing with strings
	"unicode" // to check if the letter is legit
)

// uhh,  I aint have not idea how to describe the purpose of the function, it just count words?
func countWords(text string) int {
	// If text is empty, go touch grass
	if strings.TrimSpace(text) == "" {
		return 0
	}
	words := strings.Fields(text) // Splits text into words
	return len(words)
}

// Counts how many times a word shows up, AND YEAH I DO NOT USE LOOP FOT THIS 
func countWordOccurrences(text, searchWord string) int {
	// No search word? No point in searching, right?
	if searchWord == "" {
		return 0
	}
	// Convert everything to lowercase to avoid case dramaaaaa-a-a--aa-a
	words := strings.Fields(strings.ToLower(text))
	searchWord = strings.ToLower(searchWord)
	count := 0
	for _, word := range words {
		// Skip words that don't match, keep the loop clean
		if word != searchWord {
			continue
		}
		count++ // Found a match, bump the counter
	}
	return count
}

// Finds the first word starting with a given letter, or tells you it ain't there
func findFirstWordWithLetter(text, letter string) string {
	// No letter? Read the line  12 
	if letter == "" { 
		return "No letter, no deal"
	}
	
	letter = strings.ToLower(letter)
	words := strings.Fields(text)
	for _, word := range words {
		// Check if the word starts with the letter (case-insensitive)
		if len(word) > 0 && strings.ToLower(string(word[0])) == letter {
			return word // returning words 
		}
	}
	return "No word found, try assemble" // Nothing found, bye ! 
}

func main() {
	// Set up scanner for user input, 'cause we fancy like that
	scanner := bufio.NewScanner(os.Stdin)

	// Clear the screen, make it fancy with ANSI codes
	fmt.Print("\033[2J\033[H")
	fmt.Println("\r\x1b[32m=== Yo, welcome ===\x1b[0m")
	fmt.Println("\r\x1b[34mEnter a text!\x1b[0m")
	fmt.Println("\r", "------------------------")

	for {
		// Ask for the text to analyze
		fmt.Println("\r\x1b[33m Gimme a  text:\x1b[0m")
		scanner.Scan()
		text := scanner.Text()

		// guess what ? CHECKING IF THE TEXT IS EMPTY !
		if strings.TrimSpace(text) == "" {
			fmt.Println("\r\x1b[31mError: Text is empty, your paste mode is broken ?)\x1b[0m")
			continue
		}

		// Ask for the word to search
		fmt.Println("\r\x1b[33mGimme a word to search for:\x1b[0m")
		scanner.Scan()
		searchWord := scanner.Text()

		// Print the text we're working with, make it prettyyyyyyyyyyyyyyy
		fmt.Println("\r\x1b[33mYour text is:\x1b[0m")
		fmt.Println("\r", text)
		fmt.Println("\r", "------------------------")

		// Count wods
		wordCount := countWordOccurrences(text, searchWord)
		fmt.Printf("\r\x1b[36mWord \"%s\" shows up %d times, yo!\x1b[0m\n", searchWord, wordCount)

		// Count total words
		totalWords := countWords(text)
		fmt.Printf("\r\x1b[36mTotal words in the text: %d\x1b[0m\n", totalWords)

		// Ask for the letter
		fmt.Println("\r\x1b[33mGimme a letter to find the first word with:\x1b[0m")
		scanner.Scan()
		letter := scanner.Text()

		// Check if letter is valid (just one character, and it better be a letter, or read line 12)
		if len([]rune(letter)) != 1 || !unicode.IsLetter([]rune(letter)[0]) {
			fmt.Println("\r\x1b[31mError: Letter gotta be just ONE letter, c'mon!\x1b[0m")
		} else {
			// Find the first word starting with the letter
			firstWord := findFirstWordWithLetter(text, letter)
			fmt.Printf("\r\x1b[36mFirst word starting with \"%s\": %s\x1b[0m\n", letter, firstWord)
		}

		// Simulate asking if we wanna go again ( i realized that I typed the CJ phrase from GTA only in 20 minutes..)
		fmt.Println("\r", "------------------------")
		fmt.Println("\r\x1b[35mWanna analyze another text? (yes/no):\x1b[0m")
		scanner.Scan()
		repeat := scanner.Text()

		// Using switch to handle the repeat choice
		switch strings.ToLower(repeat) {
		case "yes":
			fmt.Print("\033[2J\033[H") // Clear screen for next round
			fmt.Println("\r\x1b[32m=== Round 2, you still shouldn't touch the grass! ===\x1b[0m")
			continue
		case "no":
			fmt.Println("\r\x1b[31m Error ! Joke, program shutting down...\x1b[0m")
			break
		default:
			fmt.Println("\r\x1b[31mYo, that's not 'yes' or 'no', I'm gonna touch the grass!\x1b[0m")
			break
		}

		// Break the loop
		break
	}

	// Final cleanup, make it look clean
	fmt.Println("\r\x1b[32m=== Thanks for using the  UGA BUGA Text Analyzer! ===\x1b[0m")
}
