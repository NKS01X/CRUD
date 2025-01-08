package database

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone-no"`
	NextUser *User  `json:"next_user"`
	PrevUser *User  `json:"prev_user"`
}

// Constructor function for User struct
// returns a pointer to a User struct
func NewUser(username, email, password, phoneNo string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		PhoneNo:  phoneNo,
		NextUser: nil,
		PrevUser: nil,
	}
}

// let us imagine an head node i.e a pointer demo user
func DeleteUser(Head *User, username string) *User {
	// If the head node itself holds the username to be deleted
	if Head != nil && Head.Username == username {
		return Head.NextUser
	}

	// Search for the username to be deleted, keep track of the previous node
	// as we need to change 'prev.Next'
	prev := Head
	for prev.NextUser != nil && prev.NextUser.Username != username {
		prev = prev.NextUser
	}

	// If the username was not present in linked list
	if prev.NextUser == nil {
		return Head
	}

	// Unlink the node from linked list
	prev.NextUser = prev.NextUser.NextUser

	return Head
}
