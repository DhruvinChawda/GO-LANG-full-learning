package main

import "fmt"
import "strings"

func main(){
	var str string = "india is great country"
	str2 := []string{"india", "is", "a","great", "country"} //slice : becoz size nhi diya . size diya to array
	var result string
	var result1 string
	var result2 string
	
	result1 = strings.Replace(str, "i", "I", -1) // Replace all 'i' with 'I' in the given string because of -1
	fmt.Println("Result1 : ",result1)

	result2 = strings.Replace(str, "i", "I", 1) /*/ This will replace only first occurrence of 'i' because of 1. */
	fmt.Println("Result2: ", result2)

	result = strings.Join(str2,"_")  // Join all elements of slice with "_". Here _ will be used as separator between words.
	fmt.Println("Result : ",result)


}