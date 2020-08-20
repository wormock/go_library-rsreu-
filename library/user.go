package library

import (
	"errors"
	"strings"
)

var (
	userList = make(map[string]*User, 100)

	// ErrUserExist user already exist in db
	ErrUserExist = errors.New("user already exist")
)

// User .
type User struct {
	Login    string
	password string
	isAdmin  bool
}

// TODO Проверка что пользователь админ

func userKey(login string) string {
	return strings.TrimSpace(strings.ToLower(login))
}

// Login user
func Login(login string) *User {
	return userList[userKey(login)]
}

// AddUser to library
func AddUser(login, password string, isAdmin bool) error {
	// Check exist user
	user := Login(login)
	if user != nil {
		return ErrUserExist
	}

	// Validate
	if strings.TrimSpace(login) == "" {
		return errors.New("login is empty")
	}
	if len(password) < 3 {
		return errors.New("password must have >3 chars")
	}

	// Add user
	userList[userKey(login)] = &User{
		Login:    login,
		password: password,
		isAdmin:  isAdmin,
	}
	return nil
}
