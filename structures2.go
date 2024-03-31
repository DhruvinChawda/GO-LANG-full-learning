package main

import "fmt"

type books  struct {
	title string
	author string
	subject string
	book_id int 
}

func main(){
    var book1 books
	var book2 books

	book1.title  = "Go programming language"
	book1.author = "mohan"
	book1.subject= "Programming"
	book1.book_id=64	
	//book2 specifications
	book2.title="Python Programming"
	book2.author="mahesh"
	book2.book_id=65

	//print book1 info
	fmt.Printf("book 1 info : \n")
	fmt.Printf("Title : %s\n", book1.title)
	fmt.Printf("Author : %s \n", book1.author)
	fmt.Printf("Subject: %s \n", book1.subject)
	fmt.Printf("Book ID: %d \n", book1.book_id)

	printInfo(book1)
	printInfo(book2)

	//print book2 info
	// fmt.Printf("\nbook 2 info \n")
	// fmt.Printf("Title : %s\n", book2.title)
	// fmt.Printf("Author : %s \n", book2.author)
	// fmt.Printf("Subject: %s \n", book2.subject)
	// fmt.Printf("Book ID: %d \n", book2.book_id)
    
	
}

func  printInfo(bk books){
   fmt.Println("Title : ", bk.title)
   fmt.Println("Author : ", bk.author)
   //fmt.Println("Price : ", bk.price)
   fmt.Println("Subject : ", bk.subject)
   fmt.Println("Book ID : ", bk.book_id)
}