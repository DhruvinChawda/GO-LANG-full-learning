package main

import "fmt"

func main(){
    numbers := [8]int{10,20,30,40,50,60,70,80}
	sliceNumbers := numbers[4:7]
	fmt.Println(sliceNumbers)
	
	fmt.Printf("%T", sliceNumbers)
}