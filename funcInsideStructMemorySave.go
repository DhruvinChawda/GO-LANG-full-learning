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
	area Rectangle
	perimeter Rectangle
}


func main(){
    //assign values to struct variables

result := rectanglePara{
	length: 10,
	breadth: 20,
	color: "red",
	area: func(length int,breadth int)int {
		return length*breadth
	},

	perimeter:  func(length int, breadth int) int {
		return 2*(length+breadth)
	},
}
fmt.Println("Area of rectangle : ", result.area(result.length,result.breadth))
fmt.Println("perimeter of rectangle : ", result.perimeter(result.length,result.breadth))
fmt.Println("Color of rectangle : ", result.color)
}
