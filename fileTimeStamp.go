package main

import "fmt"
import "os"

func main(){
    myfile,err := os.Stat("sample.txt")
	if err != nil {
        fmt.Println(err)
	} else{
		fmt.Println("File modified time is : ",myfile.ModTime())
	}
}
