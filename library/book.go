package library

import (
	"errors"
	"sort"
	"strings"
)

var (
	// ErrBookExist book already exist in db
	ErrBookExist = errors.New("book already exist")
)

type sortBy int

const (
	SortDefault sortBy = iota
	SortByTitle
	SortByAuthor
	SortByYear
)

// Books library
type Books map[string]*Book

// Book .
type Book struct {
	Key              string
	Title            string
	Author           string
	YearOfPublishing int
	HoldBy           *User
}

func bookKey(key string) string {
	return strings.TrimSpace(strings.ToLower(key))
}

// Add book to library
func (b Books) Add(key, title string, author string, yearOfPublishing int) (*Book, error) {
	// 1. Проверка на дубликат (по ключу)
	if b[bookKey(key)] != nil {
		return nil, ErrBookExist
	}
	// 2. Валидация введеных параметров
	if key == "" {
		return nil, errors.New("key is empty")
	}
	if title == "" {
		return nil, errors.New("title is empty")
	}
	if author == "" {
		return nil, errors.New("author is empty")
	}
	// 3. Добавление книги
	book := &Book{
		Key:              key,
		Title:            title,
		Author:           author,
		YearOfPublishing: yearOfPublishing,
	}
	b[bookKey(key)] = book
	return book, nil
}

// SearchByTitle Поиск книги по заголовку
func (b Books) SearchByTitle(title string, sortType sortBy) ([]Book, error) {
	books := make([]Book, 0, 10)
	title = strings.TrimSpace(strings.ToLower(title))
	for _, book := range b {
		vTitle := strings.TrimSpace(strings.ToLower(book.Title))
		if strings.Contains(vTitle, title) {
			books = append(books, *book)
		}
	}
	if len(books) == 0 {
		return nil, errors.New("books with this substring not founded")
	}

	// Sort slice
	sortFunc := func(i, j int) bool { return books[i].Title < books[j].Title }

	switch sortType {
	case SortByAuthor:
		sortFunc = func(i, j int) bool { return books[i].Author < books[j].Author }
	case SortByYear:
		sortFunc = func(i, j int) bool { return books[i].YearOfPublishing < books[j].YearOfPublishing }
	}
	sort.Slice(books, sortFunc)

	return books, nil
}

// SearchByKey Поиск книги по ключу
func (b Books) SearchByKey(key string) (*Book, error) {
	book, exist := b[bookKey(key)]
	if !exist {
		return nil, errors.New("book with this key not founded")
	}
	return book, nil
}

// Hold Выдача книги пользователю
func (b Books) Hold(login string, bookKey string) (bool, error) {
	user := Login(login)
	if user == nil {
		return false, errors.New("user not founded")
	}
	book, err := b.SearchByKey(bookKey)
	if err != nil {
		return false, err
	}
	book.HoldBy = user
	return true, nil
}

// Return Возврат книги от пользователя
func (b Books) Return(login string, bookKey string) (bool, error) {
	user := Login(login)
	if user == nil {
		return false, errors.New("user not founded")
	}
	book, err := b.SearchByKey(bookKey)
	if err != nil {
		return false, err
	}
	book.HoldBy = nil
	return true, nil
}