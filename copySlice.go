package main

import "fmt"

func main(){
    primeNumbers := []int{2,3,5,7}
	numbers := []int{1,2,3,6}

	copy(numbers,primeNumbers)

	fmt.Println("Numbers : ",numbers)
}
