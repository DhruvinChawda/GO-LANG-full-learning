package main

import "fmt"

func main(){
    
	type  person struct {
		name string
		age int
	}

    person1 :=  person{name: "John", age:25 }
    fmt.Println(person1) // prints "{John 25}"

	var person2 person
	person2 = person{
		name : "zara",
		age : 29,
	}
	fmt.Println(person2) // prints "{zara 29}"
}