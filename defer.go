package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	defer fmt.Println("Goodbye!")
	defer fmt.Println("I'm going first.")
	fmt.Println("welcome to galaxy")
}