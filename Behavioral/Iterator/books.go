package main

type BookType int

const (
	Hardcover BookType = iota
	Softcover
	PaperBack
	EBook
)

type Book struct {
	Name   string
	Author string
	Pages  int
	Type   BookType
}

type Library struct {
	Collection []Book
}

func (lib *Library) IterateBooks(f func(Book) error) {
	for _, book := range lib.Collection {
		err := f(book)
		if err != nil {
			panic(err)
			break
		}
	}
}

func (lib *Library) createIterator() iterator {
	return &BookIterator{
		books: lib.Collection,
	}
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:   "War and Peace",
			Author: "Leo Tolstoy",
			Pages:  864,
			Type:   Hardcover,
		},
		{
			Name:   "Crime and Punishment",
			Author: "Leo Tolstoy",
			Pages:  1224,
			Type:   Softcover,
		},
		{
			Name:   "Brave New World",
			Author: "Aldous Huxley",
			Pages:  325,
			Type:   PaperBack,
		},
		{
			Name:   "Catcher in the Rye",
			Author: "J.D. Salinger",
			Pages:  206,
			Type:   Hardcover,
		},
		{
			Name:   "To Kill a Mockingbird",
			Author: "Harper Lee",
			Pages:  399,
			Type:   PaperBack,
		},
		{
			Name:   "Wuthering Heights",
			Author: "Emily Bronte",
			Pages:  288,
			Type:   EBook,
		},
	},
}
