package main
import("fmt")
const A int = 1
const B float32 = 3.4
var D uint8 = 200 // error will occur if value of D > 255
func main(){
	fmt.Println(A)
	fmt.Printf("Value of A: %v, Type of A: %T\n", A, A)
	fmt.Printf("Value of B: %v, Type of B: %T\n", B, B)
	//const c = 5;
	//fmt.Println(c)
	//c = 9
	//fmt.Println(c)
	fmt.Printf("Value of D: %v, Type of D: %T\n", D, D)
}