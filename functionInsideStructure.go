package main
import "fmt"
//initialize the functon rectangle
type Rectangle func(int, int) int

//create structure
type rectanglePara struct{
	length int
	breadth int 
	color string

	//function as a field of struct
	rect Rectangle
}


func main(){
    //assign values to struct variables

result1 := rectanglePara{
	length: 10,
	breadth: 20,
	color: "red",
	rect: func(length int,breadth int)int {
		return 2*(length+breadth)
	},
}

result := rectanglePara{
	length: 10,
	breadth: 20,
	color: "red",
	rect: func(length int,breadth int)int {
		return length*breadth
	},
}
fmt.Println("Area of rectangle : ", result.rect(result.length,result.breadth))
fmt.Println("perimeter of rectangle : ", result1.rect(result.length,result.breadth))
fmt.Println("Color of rectangle : ", result.color)
}
