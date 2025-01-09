package database

import (
	"errors"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone_no"`
	NextUser *User  `json:"-"`
	PrevUser *User  `json:"-"`
}

type LinkedList struct {
	Head *User
}

func (list *LinkedList) NewUser(username, email, password, phoneNo string) *User {
	newUser := &User{
		Username: username,
		Email:    email,
		Password: password,
		PhoneNo:  phoneNo,
	}

	if list.Head == nil {
		list.Head = newUser
		return newUser
	}

	current := list.Head
	for current.NextUser != nil {
		current = current.NextUser
	}
	current.NextUser = newUser
	newUser.PrevUser = current

	return newUser
}

// FindUser function searches for a user by username and returns the user pointer if present or nil if not present
func (list *LinkedList) FindUser(username string) (*User, error) {
	current := list.Head
	for current != nil {
		if current.Username == username {
			return current, nil
		}
		current = current.NextUser
	}
	return nil, errors.New("User not found")
}

// DeleteUser function deletes a user by username
func (list *LinkedList) DeleteUser(username string) {
	if list.Head == nil {
		return
	}
	if list.Head.Username == username {
		list.Head = list.Head.NextUser
		if list.Head != nil {
			list.Head.PrevUser = nil
		}
		return
	}
	current := list.Head
	for current.NextUser != nil {
		if current.NextUser.Username == username {
			current.NextUser = current.NextUser.NextUser
			if current.NextUser != nil {
				current.NextUser.PrevUser = current
			}
			return
		}
		current = current.NextUser
	}
}

// Traverse function to print all users in the linked list
// DEBUGGING PURPOSES
func (list *LinkedList) Traverse() {
	current := list.Head
	for current != nil {
		fmt.Printf("Username: %s, Email: %s, PhoneNo: %s\n", current.Username, current.Email, current.PhoneNo)
		current = current.NextUser
	}
}
