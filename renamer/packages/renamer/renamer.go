package renamer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFiles(directory, pattern string) ([]string, error) { // params are directory and a patern, it returns either string or error respectibly
	var files []string // containt for founded files
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error { // walking by directory step by step until meet a /0. Actually I hope that this function determinats an end as a /0. If not, then i don't know.

		if err != nil {
			return err
		}
		if !info.IsDir() { // if dir - skip it 
			matched, err := filepath.Match(pattern, info.Name())
			if err != nil {
				return err
			}
			if matched {
				files = append(files, path) // parsing founded files to arr
			}
		}
		return nil
	})
	return files, err // either error or files.
}


func RenameFiles(files []string, configParams []string) ([][]string, error) { // files for files, configparams for dir pattern actino etc.
	var results [][]string // results 
	scanner := bufio.NewScanner(os.Stdin) // stdin

	for _, file := range files { // iteration
		newName, err := applyRule(file, configParams) 
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", file, err)
			continue
		}

		if configParams[6] == "true" {
			fmt.Printf("Rename %s to %s? (y/n): ", file, newName) // in case you are not sure
			scanner.Scan()
			if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
				continue
			}
		}

		results = append(results, []string{file, newName})

		if configParams[5] != "true" {
			if err := os.Rename(file, newName); err != nil {
				fmt.Printf("Error renaming %s: %v\n", file, err)
				continue
			}
		}
	}
	return results, nil
}


//function for flags
func applyRule(file string, configParams []string) (string, error) {
	dir := filepath.Dir(file)
	base := filepath.Base(file)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// ahh, just options ? 
	switch configParams[2] {
	case "prefix":
		return filepath.Join(dir, configParams[3]+base), nil
	case "suffix":
		return filepath.Join(dir, name+configParams[3]+ext), nil
	case "replace":
		newName := strings.ReplaceAll(name, configParams[3], configParams[4])
		return filepath.Join(dir, newName+ext), nil
	case "extension":
		return filepath.Join(dir, name+"."+configParams[3]), nil
	case "lowercase":
		return filepath.Join(dir, strings.ToLower(base)), nil
	case "uppercase":
		return filepath.Join(dir, strings.ToUpper(base)), nil
	default:
		return "", fmt.Errorf("where : %s", configParams[2])
	}
}
