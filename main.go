package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Welcome to the Random Password Generator app!")

	low := "abcdefghijklmnopqrstuvwxyz"
	upp := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	sp := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	passwordLength := flag.Int("length", 8, "The length of the password")
	lowercase := flag.Bool("lowercase", false, "Enable use of lowercase letters in the password.")
	capital := flag.Bool("capital", false, "Enable use of capital letters in the password.")
	symbols := flag.Bool("symbols", false, "Enable use of special symbols in the password.")
	numbers := flag.Bool("numbers", false, "Enable use of numbers in the password.")

	flag.Parse()

	if *passwordLength < 8 {
		fmt.Println("The length you have specified is too small. Please select a larger length (at least 8).")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	//fmt.Println(*lowercase || *capital || *symbols || *numbers)

	if !(*lowercase || *capital || *symbols || *numbers) {
		fmt.Println("Please select at least one of the four categories: lowercase, capital, symbols or numbers.")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	//fmt.Println(*passwordLength, *lowercase, *capital, *symbols, *numbers)
	generatedPassword := ""
	var temp int
	for i := 0; i < *passwordLength; i++ {
		temp = rand.Intn(4)
		if temp == 0 {
			if *lowercase {
				generatedPassword = generatedPassword + string(low[rand.Intn(26)])
			} else {
				i = i - 1
			}
		} else if temp == 1 {
			if *capital {
				generatedPassword = generatedPassword + string(upp[rand.Intn(26)])
			} else {
				i = i - 1
			}
		} else if temp == 2 {
			if *numbers {
				generatedPassword = generatedPassword + string(num[rand.Intn(10)])
			} else {
				i = i - 1
			}
		} else if temp == 3 {
			if *symbols {
				generatedPassword = generatedPassword + string(sp[rand.Intn(27)])
			} else {
				i = i - 1
			}
		} else {
			fmt.Println("Error!")
		}
	}
	fmt.Println(generatedPassword)

}
