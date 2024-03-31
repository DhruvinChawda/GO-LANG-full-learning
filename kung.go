package main
import "fmt"

//a:=10 error dega kyuki aisa syntax globally define nahi kar sakte
var a int = 10
func main(){
	fmt.Printf("%d\n",a)
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scanf("%s",&name)
	fmt.Println("Hello",name)
	fmt.Println(name,"is","your") //error dega kiya hi print karenge usme space bar nhi mil rahegi
	}
/*	
}