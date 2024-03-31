package main

import "fmt"

func sum (nums...int){
	fmt.Println(nums," ")
	total := 0
	for _ , num := range nums{
		total+=num
	}
	fmt.Println(total)
}

func main(){
	sum(10,20)
	sum(34,56,789)
	sum(1,2,3,4,5)
	nums:=[]int{1,2,3,4}
	sum(nums...)
}