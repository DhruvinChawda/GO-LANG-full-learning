package main

import "fmt"

func main(){
	//create a map
	flowerColor := map[string]string{"sunflower":"yellow","jasmine":"white","hibiscus":"red"}
	fmt.Println(flowerColor["sunflower"])
	fmt.Println(flowerColor["jasmine"])

	//maps are mutable any values of the key can be changed

	capital := map[string]string{"Nepal":"Katmandu","Us":"new york"}
	fmt.Println("Initial map : ",capital)

	//change of us to washington dc
	capital["Us"] = "Washington D.C."
	fmt.Println(capital["Us"])
	fmt.Println("updated map : ",capital)

	students := map[int]string{1:"John",2:"Jane",3:"Tom"}
	fmt.Println("Student names: ",students)

	//add element with key 3
	students[4]="Alice"
	fmt.Println("Added student : ", students)
	delete(students,1) // delete John from the map
	fmt.Println("After deleting john : ",students)

	fmt.Println("updated map : ",students)

	squareNumber := map[int]int{2:4,3:9,4:16,5:25}
	for number ,squared := range squareNumber{
		fmt.Println(number," squared is : ",squared)
	}

	// create a map using make()
	students1:= make(map[int]string)

	//add elements to the map
	students1[1]="harry"
	students1[2]="lolo"
	students1[5]="david"

	fmt.Println(students1)

	for key := range(students1){
		fmt.Println("The key is : ",key)
	}

    
}