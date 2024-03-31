package main

import "fmt"
import "strings"

func main(){
	var str string = "Hello World"
	var subStr = "wor"

	if string.Contains(str, subStr)==true{
		fmt.Printf("String (%s) contains sub-string (%s)", str, subStr)

	}else{
		fmt.Printf("String (%s) does not contain substring (%s)", str, subStr)
	}
}