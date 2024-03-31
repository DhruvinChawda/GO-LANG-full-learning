package main

import "fmt"

func main(){

	//append function

    // primeNumbers := []int{2, 3}
	
	// primeNumbers = append(primeNumbers,5,7)
	// fmt.Println("prime numbers : ",primeNumbers)

	//program to add elements of one slice to another

	evenNumbers := []int{2,4}
	oddNumbers := []int{1,3}

	evenNumbers = append(evenNumbers,  oddNumbers...)
	fmt.Printf("%v \n", evenNumbers)
	fmt.Printf("%d",cap(evenNumbers))
}