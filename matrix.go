package main

import "fmt"

func main(){
	var sum int = 0
	var matrix1[2][2]int
	var matrix2[2][2]int	
	var matrix3[2][2]int
	
	fmt.Println("enter matrix1 elements : ")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Printf("Elements : matrix1[%d][%d]: ",i,j)
			fmt.Scan("%d",&matrix1[i][j])
		}
	}
	fmt.Println("enter matrix2 elements : ")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Printf("Elements : matrix2[%d][%d]: ",i,j)
			fmt.Scan("%d",&matrix2[i][j])
		}
	}
	fmt.Println("enter matrix3 elements : ")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Printf("Elements : matrix1[%d][%d]: ",i,j)
			fmt.Scan("%d",&matrix1[i][j])
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