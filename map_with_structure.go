package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"Bell Labs":{
		40.68433, -74.39967,
	}, 
	"Google":{
		37.42202, -122.08408,
	},
}

func main(){
    fmt.Println(m)
}
