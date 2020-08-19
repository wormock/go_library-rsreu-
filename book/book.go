package book

import (
	"errors"
	"fmt"
)

type BookList struct {
	List []*Book
}

func (l *BookList) Search(b *Book) (int, error) {
	for i, v := range l.List {
		if b.Key == v.Key || b.Title == v.Title {
			b = v
			return i, errors.New("Книга уже существует")
		}
	}
	return -1, nil
}

func (l *BookList) Append(b *Book) {
	l.List = append(l.List, b)
}

func (l *BookList) Remove(index int) {
	l.List = append(l.List[:index], l.List[index+1:]...)
}

type Writer struct {
	Name       string
	SecondName string
	Surname    string
}

type Book struct {
	Key              string
	Title            string
	Author           *Writer
	YearOfPublishing int
	IssBy            *User
}

func (b *Book) Get(u *User) error {
	if b.IssBy == nil {
		b.IssBy = u
		u.Holding = append(u.Holding, b)
	} else {
		return errors.New("Книга выдана!")
	}
	return nil
}

func (b *Book) Return(u *User) error {
	index, err := u.SearchInHolding(b)
	if err != nil {
		return err
	}
	u.RemoveFromHolding(index)
	return nil
}

type Users struct {
	Array []*User
}

func (us *Users) Search(u *User) error {
	for _, v := range us.Array {
		if v.Login == u.Login && v.password == u.password {
			u = v
			fmt.Println(u)
			fmt.Println(&v)
			return nil
		}
	}
	return errors.New("Пользователь не найден!")
}

func (us *Users) Append(u *User) {
	us.Array = append(us.Array, u)
}

type User struct {
	Login    string
	password string
	admin    bool
	Holding  []*Book
}

func (u *User) SearchInHolding(b *Book) (int, error) {
	for i, v := range u.Holding {
		if v == b {
			return i, nil
		}
	}
	return -1, errors.New("Книга не выдана этому пользователю")
}

func (u *User) RemoveFromHolding(index int) {
	u.Holding = append(u.Holding[:index], u.Holding[index+1:]...)
}
