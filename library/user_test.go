package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAdd(t *testing.T) {
	tests := []struct {
		name       string
		login      string
		password   string
		isAdmin    bool
		errMessage string
	}{
		{
			name:     "add user ok",
			login:    "evgen",
			password: "1234",
		},
		{
			name:       "add user copy error",
			login:      "evgen",
			password:   "12345",
			errMessage: "user already exist",
		},
		{
			name:       "add login empty error",
			login:      "",
			password:   "12345",
			errMessage: "login is empty",
		},
		{
			name:       "add pass <3 char error",
			login:      "evgen11",
			password:   "12",
			errMessage: "password must have >3 chars",
		},
	}

	for _, data := range tests {
		t.Run(data.name, func(tt *testing.T) {
			err := AddUser(data.login, data.password, data.isAdmin)
			if err != nil || data.errMessage != "" {
				assert.EqualError(tt, err, data.errMessage, "error not equal")
			}
		})
	}
}
