package main

import "fmt"
import "strings"

func main(){
	var str string = "India"
	var result bool = false
	var substr string = "In"

	result = strings.HasPrefix(str,substr)
	if (result == true){
		fmt.Printf("yes")
	} else{
		fmt.Println("no")
	}	
}