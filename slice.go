package main

import "fmt"

func main(){
    numbers:=[]int{1,2,3,4,5}
	fmt.Println(numbers) //[1 2 3 4 5]
	
	var numbers1 = make([]int,3,5)
	printSlice(numbers1)
}

func  printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}