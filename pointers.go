package main

import "fmt"

func main() {
	var intdata = 20

	var intpointer *int

	intpointer = &intdata

	fmt.Println("what int data stores : ", intdata)
	fmt.Println("address of intData : ", &intdata)
	fmt.Println("what  intpointer store:", intpointer)

	*intpointer = 30
	fmt.Println("what intData now stores : ", intdata)

	var name = "john"
	var ptr *string

	//assign the memory address of name to the pointer

	ptr = &name

	fmt.Println("value of pointer is ",ptr)
	fmt.Println("address of the variable ",&name)

}
