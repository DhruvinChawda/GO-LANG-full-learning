package main

import (
	"fmt"
	//"go/types"
)

//import "time"

func main(){
	// thisMonth := 5

	// switch  thisMonth {

	// case 1:
	// 	fmt.Println( "January is the first month")
	// case 2:
	// 	fmt.Println( "February is the second month.")
	// case 3:
	// 	fmt.Println("March is the third month.")
	// case 4:
	// 	fmt.Println("April is the fourth month.")
	// case 5:
	// 	fmt.Println("May is the fifth month.")
	// }

	// x:=10
	// y:=20

	// switch {
	// 	case x > y :
	// 		fmt.Printf("x is greater")
	// 	case x < y :
	// 		fmt.Println("x is lesser")
	// 	default:
	// 		fmt.Println("Both are equal")
	// }

	//time series switch case (reminder app)

	// switch today:= time.Now(); {

	// case today.Day()<5:
	// 	fmt.Println("clean your house")

	// case today.Day() <= 10:
	// 	fmt.Println("Buy some wine")

	// case today.Day()>15 && today.Day()!=25:
	// 	fmt.Println("Visit a doctor")

	// case today.Day()==25:
	// 	fmt.Println("Buy some food")

	// default:
	// 	fmt.Println("No information available")		

	// }


		//fallthrough lagane se niche next condition me jata hai
	// x:=1
	// switch x {
	// 	case 1:
	// 		fmt.Println("1")
	// 		fallthrough
	// 	case 2:
	// 		fmt.Println("2")
	// 		fallthrough
	// 	case 3:
	// 		fmt.Println("3")
	// }			

	// switch day:= 4; day{
	// case 1:
	// 	fmt.Println("Monday")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thrusday")
	// 	fallthrough
	// case 6:
	// 	fmt.Println("saturday")
	// default:
	// 	fmt.Println("Sunday")
	// }	
	
	var x interface{} = "Hello" // idhar "hello" nahi dala to type of x : <nil>
	switch i:= x.(type) {
	case nil:
		fmt.Printf("Type of x : %T",i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("X is bool or String")
	default:
		fmt.Printf("don't know the type")				

	}
			
} 