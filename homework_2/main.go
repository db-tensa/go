// YO ! So, I messed up... Recently, I reinstall my Arch Linux because of a lot of garbage. And now, I don't know WHY tcell do not display a boxes, a text and so on.
// I thought that the problem was in terminal, I switched to alacritty with zsh (shell probably didn't effect even on something), but nothing, I even asked grok, but It just said
// "Your code looks correct, try reinstall packages, or check your project " and bla bla bla, so today, just a beatiful code, and probably a  bash script for... something

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func printColor(number, numbers string) {
	const (
		Green = "\033[32m"
		Reset = "\033[0m"
	)
	first_part := number[:3] // we took the first part of code, that is 097 or 067 
	second_part := number[:6] // this is the end, it's end on 6 because program just ignore the fiveth index


	// grok generated .... basically it's just color the operator code but works weird, idk how to fix it
	fmt.Printf("\u041d\u043e\u043c\u0435\u0440: %s%s%s%s\n", first_part, Green, numbers, Reset+second_part) 
// generetated to this moment
}
func main() {

	// choise itself 
	var choise int64
	const _little_alfa string = "abcdefghijklmnopqrstuvwxyz" // ahhh idk
	const _big_alfa string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // saaame 
 

	// greeting and choises 
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("Hello Please make your choise !")

	fmt.Println("1. E-MAIL CHECk")
	fmt.Println("2. PASSWORD CHECK")
	fmt.Println("3. NUMBER CHECK")
	fmt.Println("4. IP CHECK")
	fmt.Println("5. URL CHECK")
	fmt.Println("It's workss")

	fmt.Println("----------------------------------------------------------------------------------")

	// Ok so, let's take a break, and I explay WHY bufio, but not a fmt.Scanln
	// So, go, is really really weird, I barelt starting to hate him for it
	// The buffer in go is strange unlike in C or C++
	// so after you once CALL a fmt.Scanl() in one section of the code 
	// You won't make it works again, because of the garbage in the buffer
	// So wee ned to call bufio each time when we want to call a STDIN operation
	// thanks
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// removing spaces 
	input = strings.TrimSpace(input)
	choise, _ = strconv.ParseInt(input, 10, 64)

	if choise == 1 {
		sym := "@"
		var email_string string
		fmt.Printf("Please enter an email which you would like to check >> ")
		email_string, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		email_string = strings.TrimSpace(email_string)

		// if e-mail does not contain a '@' return 1
		if strings.Count(email_string, sym) != 1 {
			fmt.Printf("Incorrect format")
			os.Exit(1)
		}

		// spliting a string
		parts := strings.Split(email_string, "@")

		// until '@'
		localPart := parts[0]
		// after '@'
		domainPart := parts[1]



		// if local part is empty return 1
		if len(localPart) == 0 {
			os.Exit(1)
		}

		// if domainPart is empty return 1
		if len(domainPart) == 0 {
			os.Exit(1)
		}


		// if whole string does not contain a '.' return 1
		if !strings.Contains(domainPart, ".") {
			os.Exit(1)
		}

		// spliting a domain part
		domainParts := strings.Split(domainPart, ".")
		tld := domainParts[len(domainParts)-1]
		if len(tld) < 2 || len(tld) > 6 {
			os.Exit(1)
		}

		if strings.Contains(email_string, " ") {
			os.Exit(1)
		}

	}

	if choise == 2 {
		// for symbols
		var counter_alfa_upper int64
		var counter_alfa_lower int64
		var counter_num int64
		var counter_special int64

		fmt.Printf("Enter a password")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		// removing spaces
		password = strings.TrimSpace(password)

		// represent a password in runes style
		runes := []rune(password)

		// getting a size of slice 
		len_password := len(runes)
		// iterate each symbol and summarize it to variabels
		for i := 0; i < len_password; i++ {
			if unicode.IsDigit(runes[i]) {
				counter_num++

			}
			if unicode.IsUpper(runes[i]) {
				counter_alfa_upper++
			}
			if unicode.IsLower(runes[i]) {
				counter_alfa_lower++
			}
			if unicode.IsSpace(runes[i]) {
				fmt.Printf("Your password contains a space ! ")
				os.Exit(1)
			}

			// this is a bad, but I don't find a function for it, or I am just blind :(
			// so basically it's just useless in most cases
			if runes[i] == '@' || runes[i] == '_' || runes [i] == '!'  {
				counter_special++
			}
		}

		// counting 
		var result int64 = counter_alfa_lower + counter_alfa_upper + counter_num

		if counter_alfa_lower == 0 {
			fmt.Print("Your password doesn't contain any  lower letter !")
		}
		if counter_alfa_upper == 0 {
			fmt.Printf("Your password doesn't contain any upper letter ! ")
		}
		if counter_num == 0 {
			fmt.Printf("Your password doesn't containt any number !")
		}
		if result < 8 {
			fmt.Printf("Your password is too short !")
		}
	}

	if choise == 3 {
		var number string
		// operators code
		valid_codes_for_sub_numbers := []string{"067", "050", "097"}


		// yes yes yes, I am a lazy person who don't care about spaces, or other  number style formats, but some websites also do not care about it. But I am still lazy... 
		fmt.Printf("Caution ! Your number contain : spaces, special symbols except '+'")
		fmt.Printf("Enter a number ! ---> ")
		number, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		number = strings.TrimSpace(number)
		len_number := len(number)

		// detecting a operator code
		sub_numbers := number[3:6]

		// WE DON'T NEED A strings.HasPrefix(str, prefix) !!!! 
		if number[0] != '+' {
			fmt.Printf("Error, the first symbol must be a '+' ! ")
			os.Exit(1)
		} else if len_number == 0 {
			fmt.Printf("Number field cannot be empty !")
			os.Exit(1)
		} else if strings.Contains(number, " ") {
			fmt.Printf("Your number field cannot contain spaces !")
			os.Exit(1)

		}

		valid_code_found := false


		// checking code operator
		for _, code_iter := range valid_codes_for_sub_numbers {
			if sub_numbers == code_iter {
				valid_code_found = true
				break
			}
		}
		// if code operator isn't in number return 1 
		if !valid_code_found {
			fmt.Printf("It seems to be your number operator is wrong ! ")
			printColor(number, sub_numbers)
			os.Exit(1)

		}
		fmt.Printf("Correct !")

	}

	if choise == 4 {
		var ip string
		fmt.Printf("Enter an IP address! ---> ")
		reader := bufio.NewReader(os.Stdin)
		ip, _ = reader.ReadString('\n')
		ip = strings.TrimSpace(ip)

		if strings.Contains(ip, " ") {
			fmt.Printf("IP cannot contain spaces!")
			os.Exit(1)
		}

		parts := strings.Split(ip, ".")
		if len(parts) != 4 {
			fmt.Printf("IP must have exactly 4 parts separated by dots!")
			os.Exit(1)
		}


		// improtant stuff
		if len(ip) < 7 || len(ip) > 15 {
			fmt.Printf("IP length must be between 7 and 15 characters!")
			os.Exit(1)
		}


		// iterating each section of ip, if some section is empty, we increment i and returning a number of empty sections
		for i, part := range parts {
			if len(part) == 0 {
				fmt.Printf("Part %d is empty!", i+1)
				os.Exit(1)
			}



			if i == 0 && len(part) > 1 && part[0] == '0' {
				fmt.Printf("First part cannot start with zero unless it's a single 0!")
				os.Exit(1)
			}

			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Part %d is not a valid number!", i+1)
				os.Exit(1)
			}


			// checking the range 
			if num < 0 || num > 255 {
				fmt.Printf("Part %d must be between 0 and 255!", i+1)
				os.Exit(1)
			}
		}


					// for UNICODE or some symbols, ip must contains only numbers from 1 to 9
		for _, r := range ip {
			if !unicode.IsDigit(r) && r != '.' {
				fmt.Printf("IP contains invalid character '%c'!", r)
				os.Exit(1)
			}
		}

		fmt.Printf("IP '%s' is valid\n", ip)
	}

	if choise == 5 {
		var url string
		fmt.Printf("Enter a URL! ---> ")
		reader := bufio.NewReader(os.Stdin)
		url, _ = reader.ReadString('\n')
		url = strings.TrimSpace(url)

		if strings.Contains(url, " ") {
			fmt.Printf("URL cannot contain spaces!")
			os.Exit(1)
		}

		// checking if string starts with http or https, if not return 1
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			fmt.Printf("URL must start with http:// or https://!")
			os.Exit(1)
		}

		
		domainStart := 7
		if strings.HasPrefix(url, "https://") {
			domainStart = 8
		}


		// start checking for protocols, initially checking a initial part of domain that is https or http
		domainEnd := len(url)
		for i, r := range url[domainStart:] {
			if r == '/' || r == '?' || r == '#' {
				domainEnd = domainStart + i
				break
			}
		}

		// checking for dots 
		domain := url[domainStart:domainEnd]
		if !strings.Contains(domain, ".") {
			fmt.Printf("Domain must contain at least one dot!")
			os.Exit(1)
		}

		// checking for UNICODE symbols
		for _, r := range domain {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.' && r != '-' {
				fmt.Printf("Domain contains invalid character '%c'!", r)
				os.Exit(1)
			}
		}

		// if you enter something which do not supposed to be in a domain a.k.a google.?com
		if domainEnd < len(url) {
			path := url[domainEnd:]
			if path[0] != '/' && path[0] != '?' && path[0] != '#' {
				fmt.Printf("Invalid character after domain!")
				os.Exit(1)
			}



			// for something not predictable...
			if strings.HasPrefix(path, "/") {
				for _, r := range path[1:] {
					if r == '?' || r == '#' {
						break
					}
					if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '/' && r != '-' && r != '_' {
						fmt.Printf("Path contains invalid character '%c'!", r)
						os.Exit(1)
					}
				}
			}


				}

		fmt.Printf("URL '%s' is valid\n", url)
	}
}
