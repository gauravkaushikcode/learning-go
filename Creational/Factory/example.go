package main

import "fmt"

func main() {
	magazine1, _ := newPublication("magazine", "Author1", 35, "Publisher1")
	magazine2, _ := newPublication("magazine", "Author2", 55, "Publisher2")
	newspaper1, _ := newPublication("newspaper", "Author1", 55, "Publisher1")
	newspaper2, _ := newPublication("newspaper", "Author2", 55, "Publisher2")

	pubDetails(magazine1)
	pubDetails(magazine2)
	pubDetails(newspaper1)
	pubDetails(newspaper2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------------------------\n")
}
