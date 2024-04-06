package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

/*

CHANGELOG
--- v0.1 ---
FEATURES:
1) Endpoint CreateUser. No soporta el servicio externo de PLD
2) Endpoint GetUserInformation. No hay auth, solo username por parametro.

TO-DO
1) consumir PLD externo
2) Endpoint login - respuesta de token
3) usar el token de respuesta como autenticacion de GetUserInformation
4) Refactor - organizar en paquetes/modularizar.

*/

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// POST /users
	// GET  /users/{username}
	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
		r.Route("/{username}", func(r chi.Router) {
			r.Use(userCtx)
			r.Get("/", getUserInformation)
		})
	})
	fmt.Println("Initializing API Rest Server on port 8000...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *User
		var err error

		if username := chi.URLParam(r, "username"); username != "" {
			user, err = dbGetUser(username)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	data := &UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	user := data.User
	dbNewUser(user)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewUserResponse(user))
}

func getUserInformation(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*User)
	if err := render.Render(w, r, NewUserResponse(user)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

type ErrResponse struct {
	Err            error  `json: "-"`
	HTTPStatusCode int    `json: "-"`
	StatusText     string `json:"status"`
	ErrorText      string `json:"error,omitempty"`
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (rd *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type UserPayload struct {
	*User
}

type UserRequest struct {
	*User
}

type User struct {
	ID        string `json: "ID"`
	Password  string `json: "password"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Email     string `json: "email"`
}

// User fixture data
var users = []*User{
	{ID: "Juanpe", Password: "1234", FirstName: "Juan", LastName: "Perez", Email: "juan@perez.com"},
	{ID: "john1234", Password: "5678", FirstName: "John", LastName: "Deere", Email: "johndeere@gmail.com"},
}

func dbGetUser(id string) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("User not found")
}

func dbNewUser(user *User) (string, error) {
	user.ID = uuid.NewString()
	users = append(users, user)
	return user.ID, nil
}

func NewUserResponse(user *User) *UserPayload {
	resp := &UserPayload{User: user}
	return resp
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.User == nil {
		return errors.New("missing required User fields")
	} else { //required fields
		if a.FirstName == "" || a.LastName == "" || a.Email == "" {
			return errors.New("missing required User fields")
		}
	}
	return nil
}
