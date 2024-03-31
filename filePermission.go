package main

import "fmt"
import "os"

func main(){
    myfile,err := os.Stat("sample.txt")
	if err != nil {
        fmt.Println(err)
	} else{
		fmt.Println("file permission : ",myfile.Mode())
	}
}

// read 4
// write 2
// executeg 1
