package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"

	clipboard "github.com/atotto/clipboard"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Println("Welcome to the Random Password Generator app!")

	low := "abcdefghijklmnopqrstuvwxyz"
	upp := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	sp := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	var passwordStorage []string
	var entropyStorage []float64

	passwordLength := flag.Int("length", 8, "The length of the password")
	lowercase := flag.Bool("lowercase", false, "Enable use of lowercase letters in the password.")
	capital := flag.Bool("capital", false, "Enable use of capital letters in the password.")
	symbols := flag.Bool("symbols", false, "Enable use of special symbols in the password.")
	numbers := flag.Bool("numbers", false, "Enable use of numbers in the password.")
	outputFile := flag.String("output", "", "Select a file to output the password to. If the file already exists, the password will be appended to the existing file.")
	noAmbiguous := flag.Bool("no-ambiguous", false, "Exclude characters that can be confusing (l,I,O,0)")
	characterPool := flag.String("character-pool", "", "Select your own character pool for the password generation.")
	count := flag.Int("count", 1, "Select the count of passwords to be generated.")
	copyToClipboard := flag.Bool("clipboard", false, "Copy the password on the system clipboard after creating. If you select count>1, the last password generated will be copied to clipboard.")

	flag.Parse()

	if *passwordLength < 4 {
		fmt.Println("The length you have specified is too small. Please select a larger length (at least 4).")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	} else if *passwordLength > 256 {
		fmt.Println("The length you have specified is too large. Please select a smaller length (maximum 256).")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	if (*characterPool == "") && !(*lowercase || *capital || *symbols || *numbers) {
		fmt.Println("Please select at least one of the four categories: lowercase, capital, symbols or numbers.")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	if ((*characterPool != "") && (*lowercase || *capital || *symbols || *numbers)) {
		fmt.Println("You can either select a custom character pool or one of the four options (lowercase, capital, symbols or numbers).")
		fmt.Println("Program Exiting.")
		os.Exit(0)
	}

	if *noAmbiguous {
		low = "abcdefghijkmnopqrstuvwxyz"
		upp = "ABCDEFGHJKLMNPQRSTUVWXYZ"
		num = "123456789"
	}

	var pool string

	if *characterPool != "" {
		pool = *characterPool
	} else {
		pool = ""
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
	}

	//var nbig *big.Int
	//var n int
	//var mylen *big.Int
	
	for j := 0; j < *count; j++ {
		generatedPassword := ""
		for i := 0; i < *passwordLength; i++ {
			// generatedPassword = generatedPassword + string(pool[ rand.Intn(len(pool)) ])
			//mylen :=  big.NewInt(int64(len(pool)))
			nbig, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
			if err != nil {
				panic(err)
			}
			n := nbig.Int64()
			generatedPassword = generatedPassword + string(pool[n])
		}
		passwordStorage = append(passwordStorage, generatedPassword)
		fmt.Println(passwordStorage[j])
		
		passwordEntropy := passwordvalidator.GetEntropy(passwordStorage[j])
		entropyStorage = append(entropyStorage, passwordEntropy)
		// fmt.Println(entropyStorage[j])
		fmt.Print("Password Strength: ")
		if entropyStorage[j] < 28 {
			fmt.Println("Weak")
		} else if entropyStorage[j] < 35 {
			fmt.Println("Moderate")
		} else if entropyStorage[j] < 59 {
			fmt.Println("Strong")
		} else {
			fmt.Println("Very Strong")
		}
		if *outputFile != "" {
			f, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}
			if _, err = f.WriteString(passwordStorage[j]+"\n"); err != nil {
				panic(err)
			}
			defer f.Close()
			os.Executable()
		}
		if *copyToClipboard {
			clipboard.WriteAll(passwordStorage[j])
		}
	}
	
}
