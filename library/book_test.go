package library

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookAdd(t *testing.T) {
	tests := []struct {
		name             string
		key              string
		title            string
		author           string
		yearOfPublishing int
		holdBy           *User
		errMessage       string
	}{
		{
			name:             "add book ok",
			key:              "32d",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "",
		},
		{
			name:             "add book error",
			key:              "32d",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "book already exist",
		},
		{
			name:             "add book error",
			key:              "",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "key is empty",
		},
		{
			name:             "add book error",
			key:              "33d",
			title:            "",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "title is empty",
		},
		{
			name:             "add book error",
			key:              "34d",
			title:            "Война и мир",
			author:           "",
			yearOfPublishing: 2020,
			errMessage:       "author is empty",
		},
	}
	books := make(Books, len(tests))
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.Add(data.key, data.title, data.author, data.yearOfPublishing)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}

func TestBookSearchByKey(t *testing.T) {
	tests := []struct {
		name             string
		key              string
		title            string
		author           string
		yearOfPublishing int
		holdBy           *User
		errMessage       string
	}{
		{
			name:             "search by key ok",
			key:              "32d",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "",
		},
		{
			name:             "search by key error",
			key:              "33ddd",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "book with this key not founded",
		},
	}
	books := make(Books, len(tests))
	_, err := books.Add("32d", "Война и мир", "Лев Николаевич Толстой", 2020)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.SearchByKey(data.key)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}

func TestBookSearchByTitle(t *testing.T) {
	tests := []struct {
		name             string
		key              string
		title            string
		author           string
		yearOfPublishing int
		holdBy           *User
		errMessage       string
	}{
		{
			name:             "search by title ok",
			key:              "32d",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2019,
			errMessage:       "",
		},
		{
			name:             "search by title ok",
			key:              "33d",
			title:            "Война миров",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2022,
			errMessage:       "",
		},
		{
			name:             "search by title ok",
			key:              "34d",
			title:            "Мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 1999,
			errMessage:       "",
		},
		{
			name:             "search by title eror",
			key:              "36d",
			title:            "Чайка",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 1830,
			errMessage:       "books with this substring not founded",
		},
	}
	books := make(Books, len(tests))
	_, err := books.Add("32d", "Война и мир", "Лев Николаевич Толстой", 2020)
	_, err = books.Add("33d", "Война миров", "Лев Николаевич Толстой", 2020)
	_, err = books.Add("34d", "Мир", "Лев Николаевич Толстой", 2020)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.SearchByTitle(data.title, SortByTitle)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
			_, err = books.SearchByTitle(data.title, SortByAuthor)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
			_, err = books.SearchByTitle(data.title, SortByYear)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}

func TestBookHold(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		login      string
		errMessage string
	}{
		{
			name:       "hold ok",
			key:        "32d",
			login:      "wormock",
			errMessage: "",
		},
		{
			name:       "hold error",
			key:        "33ddd",
			login:      "wormock",
			errMessage: "book with this key not founded",
		},
		{
			name:       "hold error",
			key:        "32d",
			login:      "krutoi_poc",
			errMessage: "user not founded",
		},
	}
	books := make(Books, len(tests))
	_ = AddUser("wormock", "admin", true)
	_, err := books.Add("32d", "Война и мир", "Лев Николаевич Толстой", 2020)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.Hold(data.login, data.key)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}

func TestBookReturn(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		login      string
		errMessage string
	}{
		{
			name:       "return ok",
			key:        "32d",
			login:      "wormock",
			errMessage: "",
		},
		{
			name:       "return error",
			key:        "33ddd",
			login:      "wormock",
			errMessage: "book with this key not founded",
		},
		{
			name:       "return error",
			key:        "32d",
			login:      "molodoy_platon",
			errMessage: "user not founded",
		},
		{
			name:       "return error",
			key:        "32d",
			login:      "krutoi_poc",
			errMessage: "book was given to another user",
		},
	}
	books := make(Books, len(tests))
	_ = AddUser("wormock", "admin", true)
	_ = AddUser("krutoi_poc", "neAdmin", false)
	_, err := books.Add("32d", "Война и мир", "Лев Николаевич Толстой", 2020)
	_, err = books.Hold("wormock", "32d")
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.Return(data.login, data.key)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}
