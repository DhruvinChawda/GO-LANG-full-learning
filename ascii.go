package main

import "fmt"

func main(){
	str:="Hello World!"

	fmt.Println("Ascii value of string");
	for i:=0; i<len(str); i++{
		fmt.Printf("%c %d\n",str[i],str[i])
	}
}