package main

import (
	"fmt"
	"sort"
)

func main(){
    slice :=  []int{70,20,30,40,50,60,10,80,90,100}

	sort.Ints(slice) //Sorting the Slice in Ascending Order

	fmt.Println( "The sorted array is: ", slice )
	for _,ele := range slice {
		fmt.Printf("%d -> ",ele)
	}
	// Sorting in descending order using sort.Sort with sort.Reverse
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	fmt.Println("\nSorted in descending order:", slice)

	slice1 := []string{"honesty","is","the", "best","policy"}
	sort.Strings(slice1)
	fmt.Println("Sorted words are :")
	for _,item := range slice1 {
		fmt.Printf("%s ",item)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(slice1)))

	fmt.Println("\nSorted in descending order:", slice1)
}