package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is cipher")
var secret = flag.String("secret", "", "Your secret key. Must be contain at least 1 character")

func main() {
	flag.Parse()

	if len(*secret) == 0 {
		fmt.Fprintf(os.Stderr, "Please provide your secret")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plainText := getUserInput("Enter your text to cipher:")
		fmt.Println(plainText)

		cipheredText, err := cipherer.Cipher(plainText, *secret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypted text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("Enter your ciphered text to decipher:")
		fmt.Println(cipheredText)

		decipheredText, err := cipherer.Decipher(cipheredText, *secret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decrypted text; %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Println("Invalid mode. Exiting now...")
		os.Exit(1)
	}
}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err, "Please Try again")
			continue
		}

		return strings.TrimRight(result, "\n")
	}
}
