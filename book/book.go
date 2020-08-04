package book

import (
	"errors"
)

type BookList struct {
	List []Book
}

func (l *BookList) SearchBook(b Book) (Book, int, error) {
	for i, v := range l.List {
		if b.Key == v.Key || b.Title == v.Title {
			return v, i, errors.New("Книга уже существует")
		}
	}
	return b, -1, nil
}

func (l *BookList) AppendBook(b Book) {
	l.List = append(l.List, b)
}

type Writer struct {
	Name       string
	SecondName string
	Surname    string
}

type Book struct {
	Key              string
	Title            string
	Author           Writer
	YearOfPublishing int
	IssBy            string
}
