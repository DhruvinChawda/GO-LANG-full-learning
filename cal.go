package main

import(
	"fmt"
	"packages/calculator"

)

func main(){
	number1 := 9
	number2 := 5

	fmt.Println(calculator.Add(number1, number2)) // Output:  14
	fmt.Println(calculator.Subtract(number1, number2))// Output:  4

}