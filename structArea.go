package main

import "fmt"

func main(){
    type rectangle struct{
		length int
		breadth int
	}

	rect := rectangle{22,12}
	fmt.Println("Length : ",rect.length)
	fmt.Println("Breadth: ", rect.breadth)

	area := rect.length * rect.breadth
	perimeter:= 2*(rect.length + rect.breadth)
	fmt.Printf("\nArea of  the Rectangle is %d\n", area)
	fmt.Printf("Perimeter of  the Rectangle is %d\n", perimeter)
}
