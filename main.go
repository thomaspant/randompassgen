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
	} else if *passwordLength > 256 {
		fmt.Println("The length you have specified is too large. Please select a smaller length (maximum 256).")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	//fmt.Println(*lowercase || *capital || *symbols || *numbers)

	if !(*lowercase || *capital || *symbols || *numbers) {
		fmt.Println("Please select at least one of the four categories: lowercase, capital, symbols or numbers.")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}
	pool := ""
	if *lowercase {
		pool = pool + low
	}
	if *capital {
		pool = pool + upp
	}
	if *symbols {
		pool = pool + sp
	}
	if *numbers {
		pool = pool + num
	}

	//fmt.Println(*passwordLength, *lowercase, *capital, *symbols, *numbers)
	generatedPassword := ""
	for i := 0; i < *passwordLength; i++ {
		generatedPassword = generatedPassword + string(pool[rand.Intn(len(pool))])
	}
	fmt.Println(generatedPassword)

}
