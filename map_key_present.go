package main

import "fmt"

func main(){
    m := make(map[string]int)

	m["Answer"]=42
	fmt.Println("The value ",m["Answer"]) // 42

	m["Answer"]=48
	fmt.Println("The value ",m["Answer"]) //48

	delete(m,"Answer")
	fmt.Println(m["Answer"]," is deleted.")//0 is deleted.

	v, ok := m["Answer"]
	fmt.Println("the value : ", v, "\t exist? ",ok)//0 false


}
