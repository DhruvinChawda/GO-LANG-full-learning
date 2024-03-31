package main
import "fmt"

func main(){
	// var x float64 = 20.69
	// y:=42
	// //p:=78
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Printf("Value of x is %f and x is of type : %T\n",x,x)
	// fmt.Printf("Value of y is %d and y is of type : %T\n",y,y)
	// print("Hello world\n")
	// print("jello world")
	// var a int = 5
	// var b float64
	// b = float64(a)
	// print("value b = %f",b)
	//fmt.Printf(p)

	var(
		a int
		b int = 1
		c string = "hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%v \t%T\n", c, c)
	fmt.Printf("%v \t%T\n", a, a)
	fmt.Printf("%v \t%T\n", b, b)
}




}
