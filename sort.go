package main

import "fmt"
import "sort"

func main(){
	numbers := []int{4,2,7,1,9,3}
	numbers2 := []float64{4.1,2.3,7.4,1.2,9.4,3.5,3.6}
	fruits := []string{"banana","apple","orange","grape","pineapple"}

	sort.Ints(numbers)
	sort.Float64s(numbers2)
	sort.Strings(fruits)

	fmt.Println(numbers) // [1 2 3  4 7 9]
	fmt.Println(numbers2)
	fmt.Println(fruits) // ["apple banana grape orange pineapple"]
}