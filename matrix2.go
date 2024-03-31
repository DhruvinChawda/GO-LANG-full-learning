package main

import (
	"fmt"
)

// reverseNumber takes an integer and returns its reverse
func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		digit := num % 10        // Extract the last digit of num
		reversed = reversed*10 + digit // Append digit to reversed
		num /= 10               // Remove the last digit from num
	}
	return reversed
}

func main() {
	var number int
	fmt.Print("Enter a number to reverse: ")
	fmt.Scanln(&number)

	reversedNumber := reverseNumber(number)
	fmt.Printf("Reversed Number: %d\n", reversedNumber)
}
