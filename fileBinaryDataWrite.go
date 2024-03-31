package main

import "encoding/binary"
import "fmt"
import "os"

type Str struct{
	intNum uint8
	floatNum float32
}

func main(){
	file,err := os.Create("data.bin")
	if err != nil {
		panic("Could not create file")
	}
	var st = Str{10, 2.3}

	err = binary.Write(file, binary.LittleEndian, st)
	if err != nil {
		panic("Failed to write data to file")
	}
	fmt.Println("structure written into file successfully")

	file.Close()
}