package main

type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) hasNext() bool {
	return b.current < len(b.books)
}

func (b *BookIterator) next() *Book {
	if b.hasNext() {
		book := b.books[b.current]
		b.current++
		return &book
	}
	return nil
}
