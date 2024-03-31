package main

import "fmt"
import "os"

func main(){
    file,err := os.Create("sample.txt")
	if err != nil{
        fmt.Println("Error creating the file: %s", err)
	}

	len,err := file.WriteString("Hello world!")
	if  err != nil {
        fmt.Printf("Error writing to file: %s", err)
	}
	file.Close()
    
    // Read and print the content of 'sample.txt'
    contents, _ := os.ReadFile("sample.txt")
    fmt.Println(string(contents))
	fmt.Println(len,"bytes written to sample.txt.")
}


//textData, err = ioutil.ReadFile("sample.txt")
//fmt.printf("%s",textdata)
	
    

