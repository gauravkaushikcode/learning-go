package main

import "fmt"

func main() {
	//lib.IterateBooks(myBookCallback)
	//
	//lib.IterateBooks(func(book Book) error {
	//	fmt.Println("Book author:", book.Author)
	//	return nil
	//})

	iterator := lib.createIterator()
	for iterator.hasNext() {
		fmt.Printf("Book %+v\n", iterator.next())
	}
}

func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
