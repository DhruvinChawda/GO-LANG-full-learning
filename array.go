package main

import "fmt"

func main(){
	var arraynum = [5]int{1,5,8,2,7}
	fmt.Println(arraynum) //[1 5 8 2 7]

	var arraynum2[3] int
	arraynum2[0]=5
	arraynum2[1]=10
	arraynum2[2]=15
	fmt.Println(arraynum2)

	arraynum3 := [5]int{0:7,3:9}
	fmt.Println(arraynum3)//[7 0 0 9 0]

	var arraynum4 = [...]string{"Hello","Rajesh"}
	fmt.Println(arraynum4)//[Hello Rajesh] [...] for dynamic size
}