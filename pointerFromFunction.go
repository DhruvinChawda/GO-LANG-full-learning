package main

import "fmt"

func main(){
    result := display()
	fmt.Println("welcome to ",*result)
}

func display()*string{
   var str *string = new(string)  //declare a pointer variable of type string and initialize it with nil value
    str1 := "Hello World!"  // string literal
    return &str1             // pointer to the string literal 
}