package main

import "fmt"
import "os"
import "io"

func main(){
    existingFile, err := os.Open("existing_file.txt") // For read access.
	if err != nil {                            // Check if error occurred while opening the file.
        fmt.Println("Error: %v", err)           // If yes, print an error message and terminate the program.
        return                                // Return from the function to avoid further execution.
    }

	CopyFile, err := os.Create("copy_of_file.txt")  // Create a new file for writing.
	if err != nil{
		fmt.Println("Unable to create file")
	}
	len,err := io.Copy(CopyFile, existingFile)      // Copy the contents of the source file to the destination file.
	if err!=nil{
		fmt.Println("unable to copy file")
	}
    fmt.Printf("%d bytes copied",len)
    existingFile.Close()
	CopyFile.Close()
}
