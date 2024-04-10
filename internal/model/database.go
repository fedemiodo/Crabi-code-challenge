package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var users = []*User{}

func DbGetUserFromCredentials(user *User) (*User, error) {
	time.Sleep(100 * time.Millisecond)
	for _, u := range users {
		if user.Email == u.Email && user.Password == u.Password {
			return u, nil
		}
	}
	return nil, errors.New("invalid User/Pass")
}

func DbNewUser(user *User) error {
	// Primero busco que no se repita
	if existingUser, _ := DbGetUserFromCredentials(user); existingUser != nil {
		// encontre usuario existente
		return errors.New("user already registered")
	} else {
		time.Sleep(100 * time.Millisecond)
		user.AccessToken = &AccessToken{AccessToken: ""}
		users = append(users, user)
		return nil
	}

}

func DbGetUserByToken(token *AccessToken) (*User, error) {
	time.Sleep(100 * time.Millisecond)
	for _, u := range users {
		if u.AccessToken.AccessToken == token.AccessToken {
			return u, nil
		}
	}
	return nil, errors.New("Unauthorized")
}
func DbUpdateUserAccessToken(user *User) (*User, error) {
	// busco objeto
	if u, err := DbGetUserFromCredentials(user); err != nil {
		return nil, err
	} else {
		u.AccessToken = &AccessToken{AccessToken: uuid.NewString()}
		return u, nil
	}
}
