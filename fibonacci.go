package main

import "fmt"

func main(){
	var num1 int = 0  
	var num2 int = 1
	var num3 int = 0
	//var num4 int = 3

	fmt.Printf("%d ",num1)
	fmt.Printf("%d ",num2)

	for i:=0;i<10;i++{
		num3 = num1 + num2
		fmt.Printf( "%d ",num3 )
		num1 = num2
		num2 = num3
	}

// repeat:
	
// 	num3 = num1 + num2
// 	fmt.Printf( "%d ",num3)
// 	num1 = num2
// 	num2 = num3

// 	num4 = num4 + 1
// 	if num4 <= 10{
// 		goto repeat
// 	}
}