package main

import "fmt"
import "os"

func main(){
    programName := os.Args[0]
	// arg1 := os.Args[1]
	// arg2 := os.Args[2]
	
	fmt.Println("Total arguments :",len(os.Args))
	fmt.Println("program Name",programName)

	fmt.Println(programName)
	if len(os.Args) == 1 {
        fmt.Println("No arguments provided.")
        return
    } else if len(os.Args) > 3 {
    	fmt.Printf("%s takes exactly two arguments.\n", programName)
    	return
    }
}
