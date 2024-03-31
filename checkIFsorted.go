package main

import "fmt"
import "sort"

func main(){
    
	var status bool = false

	slice1 := []int{10,20,30,40,50,60,70,80,90,100}
	slice2 := []int{30,60,20,10,90,80,100,70,40,50} 
	slice3 := []string{"my","name","is","kakashi"}
	slice4 := []float64{1.2,4.5,6.5} 

	status = sort.IntsAreSorted(slice1)  //Checking if slice1 is sorted or not
	if status == true{
		fmt.Println( "Slice 1 is already sorted")
	}else {
		fmt.Println("Slice 1 is not sorted")
	}

	status = sort.IntsAreSorted(slice2)
	if status == true{
		fmt.Println( "Slice 2 is already sorted")
	}else {
		fmt.Println("Slice 2 is not sorted")
	}

	status = sort.StringsAreSorted(slice3)
	if status == true{
		fmt.Println( "Slice 3 is already sorted")
	}else {
		fmt.Println("Slice 3 is not sorted")
	}

	status = sort.Float64sAreSorted(slice4)
	if status == true{
		fmt.Println( "Slice 4 is already sorted")
	}else {
		fmt.Println("Slice 4 is not sorted")
	}
}