package main

import "fmt"
import "encoding/json"

func main(){
    colors := []string{"red","green","blue"}
	result,_ := json.Marshal(colors)
	fmt.Println(string(result))
	fmt.Println(result)
}
