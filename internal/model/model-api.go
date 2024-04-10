package model

import (
	"errors"
	"net/http"
)

// Required api behavior de USERS - Acoplamiento - A mejorar

type UserPayload struct {
	*User
}

type UserRequest struct {
	*User
}

type UserLoginRequest struct {
	*User
}

func (rd *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.User == nil {
		return errors.New("missing required User fields")
	} else {
		if err := validateRequiredUserFields(a); err != nil {
			return err
		}
		return nil
	}
}

func (a *UserLoginRequest) Bind(r *http.Request) error {
	if a.User == nil {
		return errors.New("missing required User fields")
	} else {
		if err := validateRequiredUserLoginFields(a); err != nil {
			return err
		}
		return nil
	}
}

func validateRequiredUserFields(a *UserRequest) error {
	switch {
	case a.FirstName == "":
		return errors.New("missing field firstName (string)")
	case a.LastName == "":
		return errors.New("missing field lastName (string)")
	case a.Email == "":
		return errors.New("missing field email (string)")
	case a.Password == "":
		return errors.New("missing field password (string)")
	default:
		return nil
	}
}

func validateRequiredUserLoginFields(a *UserLoginRequest) error {
	switch {
	case a.Email == "":
		return errors.New("missing field email (string)")
	case a.Password == "":
		return errors.New("missing field password (string)")
	default:
		return nil
	}
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
}

func (token *AccessToken) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
