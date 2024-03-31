package main

import "fmt"

func main(){
	fmt.Print("Enter first String: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter second string")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first+" "+second)
}