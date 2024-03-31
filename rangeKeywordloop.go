package main

import "fmt"

func main(){
	colors:=[]string{"Red", "Blue", "green"}
	for index,val:= range colors{
		fmt.Println(index,"=>", val)
	}
	for _,val:= range colors{
		fmt.Println(val)
	}
	for index,_:= range colors{
		fmt.Println(index)
	}

}
