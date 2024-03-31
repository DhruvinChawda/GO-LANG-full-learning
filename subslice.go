package main

import "fmt"

func main(){
	//create a slice
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	//print the original slice
	fmt.Println("Original slice:",numbers)

	fmt.Println("numbers[1:4] : ", numbers[1:4]) //get elements from index 1 to 3 (not included  3)
	fmt.Println("numbers[:3] : ", numbers[:3])     //from first till index 3 (included 3)
	fmt.Println("numbers[4:] : ", numbers[4:])    //from index 3 till the end of the array

	numbers1 := make([]int,0,5)                      //make creates an empty slice with length 5 and capacity 5
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)

}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}