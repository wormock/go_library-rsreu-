package library

import (
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
		{
			name:             "search by key ok",
			key:              "3233d",
			title:            "Война и мир",
			author:           "Лев Николаевич Толстой",
			yearOfPublishing: 2020,
			errMessage:       "",
		},
		// {
		// 	name:             "search by key error",
		// 	key:              "33ddd",
		// 	title:            "Война и мир",
		// 	author:           "Лев Николаевич Толстой",
		// 	yearOfPublishing: 2020,
		// 	errMessage:       "book with this key not founded",
		// },
	}
	books := make(Books, len(tests))
	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			_, err := books.Add(data.key, data.title, data.author, data.yearOfPublishing)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
			_, er := books.SearchByKey(data.key)
			if er != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}
