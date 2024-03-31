package main

import "fmt"

func main() {
	i := 0
loop:
	if i <= 4 {
		fmt.Println(i)
		i++
		goto loop
	}
}
