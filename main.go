package main

import (
	"fmt"
	"go_library-rsreu-/library"
)

func main() {
	bList := make(library.Books, 100)
	books, err := bList.SearchByTitle("Царевна", library.SortByAuthor)
	fmt.Println(books, err)
}
