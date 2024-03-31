package main

import "fmt"

func main(){
    // numbers := []int{2,4,6,8,10}
	// numbers2 := []int{1,3,5,7,9}

	// for i:=0;i<len(numbers);i++ {
    //     fmt.Println("The value of number at index",i,"is: ",numbers[i])
	// }

	arr := [10]int{10,20,30,40,50,60,70,80,90,100}

	intSlice := arr[2:5]

	fmt.Println("slicing elements : ")
	for index, ele := range intSlice{
		fmt.Printf("index = %d, element = %d\n",index,ele)
	}

	for _, ele := range intSlice{
		fmt.Printf("element = %d\n",ele)
	}

	for index, _ := range intSlice{
		fmt.Println(intSlice[index])
	}

	intSlice[1]=500
	fmt.Println("\nafter modifying the slice")
	for index, ele := range intSlice{
		fmt.Printf("index = %d, element = %d\n",index,ele)
	}

	newSlice := make([]int, 3, 5)
	newSlice[0], newSlice[1], newSlice[2] = 1, 2, 3
	fmt.Println("\ncreated a new slice using make function and assigned values to it.")
	fmt.Println("Elements in the newly created slice are as follows:")
	for _, val := range newSlice{
		fmt.Printf("%d ",val)
	}
}
