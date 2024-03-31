package main

import "fmt"

func main(){
	var sum int = 0
	var matrix[3][3]int
	//matrix = [3][3]int
	
	fmt.Println("enter matrix elements : ")
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Printf("Elements : matrix[%d][%d]: ",i,j)
			fmt.Scan("%d",&matrix[i][j])
		}
	}

fmt.Printf("Matrix: \n")
for i:=0;i<3;i++ {
	for j:=0;j<3;j++ {
		if i==j {
			sum = sum + matrix[i][j]
		}
		fmt.Printf("%d",matrix[i][j])
	}
	fmt.Printf("\n")
}
fmt.Printf("\nsum of diagonal elements is %d",sum)
}

