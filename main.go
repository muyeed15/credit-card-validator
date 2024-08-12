package main

import (
	"credit_card_validation_system/validation"
	"fmt"
)

func main() {
	fmt.Println(" <--------- Simple Credit Card Validation System (Luhn Algorithm) --------->")
	fmt.Print("\nEnter your credit card number: ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("\nYour card is %v!", validation.CheckLuhn(input))
}
