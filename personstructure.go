package main

import "fmt"

type person  struct {
	name string
	age int
	job string
	salary int 
}

func main(){
    var person1 person
	var person2 person

	person1.name  = "Disha"
	person1.age = 21
	person1.job= "teacher"
	person1.salary=6000	

	//person2 specifications

	person2.name="Naruto"
	person2.age=18
	person2.job="Uzumaki clan ninja"
	person2.salary=10000

	//print person1 info
	fmt.Printf("person 1 info : \n")
	fmt.Printf("name : %s\n", person1.name)
	fmt.Printf("Age : %d \n", person1.age)
	fmt.Printf("job : %s \n", person1.job)
	fmt.Printf("salary : %d \n", person1.salary)

	//print person2 info
	fmt.Printf("\nperson 1 info : \n")
	fmt.Printf("name : %s\n", person2.name)
	fmt.Printf("Age : %d \n", person2.age)
	fmt.Printf("job : %s \n", person2.job)
	fmt.Printf("salary : %d \n", person2.salary)
    
	
}