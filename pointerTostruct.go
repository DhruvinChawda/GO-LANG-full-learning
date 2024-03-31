// package main

// import "fmt"

// type vertex struct {
// 	x int
// 	y int
// }

// func main(){
//     v := vertex{1,2}
// 	p := &v
// 	(*p).x = 1e9
// 	fmt.Println(v) // {100000001 2}
// 	fmt.Println(p) // &{100000001 2}
// }

package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main(){
    v := vertex{1,2}
	p := &v
	p.x = 1e9
	fmt.Println(v) // {100000001 2}
	fmt.Println(p) // &{100000001 2}
}

