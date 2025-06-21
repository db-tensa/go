package main

import (
	"flag"
	"fmt"
	"os"

	"rename/packages/renamer"
)

func main() {
	fmt.Println("--Ill Rename--") 

	// Define flags
	dir := flag.String("dir", "./test_files/", "Directory to rename files in")
	pattern := flag.String("pattern", "*", "File pattern for example .jpg .png .cpp")
	prefix := flag.String("p", "", "Add pref")
	suffix := flag.String("s", "", "Add suff")
	replace := flag.String("r", "", "String to replace in filenames")
	replaceWith := flag.String("w", "", "Replacement string for -r")
	extension := flag.String("e", "", "New extension without dot ! ")
	lowercase := flag.Bool("l", false, "Convert filenames to lowercase")
	uppercase := flag.Bool("u", false, "Convert filenames to uppercase")
	dryRun := flag.Bool("d", false, "Enable dry run mode")
	interactive := flag.Bool("i", false, "Enable interactive confirmation")
	flag.Parse()


	// only one prefix,
	actionCount := 0
	var action, value, replaceWithValue string
	if *prefix != "" {
		action, value = "prefix", *prefix
		actionCount++
	}
	if *suffix != "" {
		action, value = "suffix", *suffix
		actionCount++
	}

	// i know i know that instead of exiting i should implement it as a recursion call of function.
	if *replace != "" {
		if *replaceWith == "" {
			fmt.Println("Error: -w required with -r")
			os.Exit(1)
		}

		// replacing files
		action, value, replaceWithValue = "replace", *replace, *replaceWith
		actionCount++
	}
	if *extension != "" {
		action, value = "extension", *extension
		actionCount++
	}
	if *lowercase {
		action = "lowercase"
		actionCount++
	}
	if *uppercase {
		action = "uppercase"
		actionCount++
	}

	if actionCount != 1 {
		fmt.Println("Error: Exactly one action flag (-p, -s, -r, -e, -l, or -u) must be provided")
		flag.Usage()
		os.Exit(1)
	}

	// preparing config 
	configParams := []string{
		*dir,
		*pattern,
		action,
		value,
		replaceWithValue,
		fmt.Sprintf("%t", *dryRun),
		fmt.Sprintf("%t", *interactive),
	}

	files, err := renamer.FindFiles(*dir, *pattern)
	if err != nil {
		fmt.Printf("Error finding files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nFound %d files\n\n", len(files))

	results, err := renamer.RenameFiles(files, configParams)
	if err != nil {
		fmt.Printf("Error during renaming: %v\n", err)
		os.Exit(1)
	}

	for _, result := range results {
		fmt.Printf("%s -> %s\n", result[0], result[1])
	}

	fmt.Printf("\nSuccessfully renamed: %d files\n", len(results))
}
