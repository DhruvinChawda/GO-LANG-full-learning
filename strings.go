package main

import "fmt"
import "strings"

func main(){
	var str string = "Hello World"
	var result string
	var result2 string
	var ind int = 0
	var str2 string = "India-is-a-great-country"
	var strArr []string = strings.Split(str2,"-")
	var str3 string = "star "
	var result3 string

	result = strings.ToUpper(str)
	result2 = strings.ToLower(str)
	ind = strings.Index(str, "W")
	result3 = strings.Repeat(str3,4)


	fmt.Println("String in Upper Case: ", result)
	fmt.Println("String in Lower Case: ", result2)
	fmt.Println("Position of 'W' : ", ind)
	fmt.Println("splitted string : ")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("%s \n", strArr[i])
	}
	fmt.Println("repeated string : ", result3)

	for i:=0; i<len(result3);i++{
		fmt.Printf("%c ",result3[i])
	}
}