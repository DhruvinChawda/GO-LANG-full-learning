package main

import "fmt"

func main(){
    places := map[string]string{"nepal":"kathmandu","us":"washington"}
	for _,capital := range places{
		fmt.Println(capital)
	}
}
