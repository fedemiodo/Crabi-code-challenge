package api

import (
	"errors"
	"net/http"

	"github.com/fedemiodo/Crabi-code-challenge/internal/model"
	"github.com/fedemiodo/Crabi-code-challenge/internal/providers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Handler(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.AllowContentType("application/json"))
		r.Post("/login", requestLogin)
		r.Post("/createNewUser", requestCreateNewUser)
		r.Get("/getMe", requestGetMe)
	})
}

func requestCreateNewUser(w http.ResponseWriter, r *http.Request) {
	data := &model.UserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	user := data.User
	if isUserBlacklisted, err := providers.ConsumePLDService(user); err != nil {
		// external API call error
		render.Render(w, r, ErrProvider(err))
	} else {
		switch isUserBlacklisted {
		case true:
			render.Status(r, http.StatusBadRequest)
			render.Render(w, r, ErrUnAuthorizedAction(errors.New("cannot create user not complying to PLD standards")))
		//render.Render(w, r, ErrInvalidRequest(err))
		case false:
			model.DbNewUser(user)
			render.Status(r, http.StatusCreated)
			if err := render.Render(w, r, UserResponse(user)); err != nil {
				render.Render(w, r, ErrRender(err))
				return
			}
		}
	}
}

func requestGetMe(w http.ResponseWriter, r *http.Request) {
	token := &model.AccessToken{AccessToken: r.Header.Get("Authorization")}
	if token.AccessToken == "" {
		// No Auth Header
		render.Status(r, http.StatusForbidden)
		render.Render(w, r, ErrInvalidRequest(errors.New("missing Authorization header")))
		return
	}
	if user, err := model.DbGetUserByToken(token); err != nil {
		// Token not found
		render.Status(r, http.StatusForbidden)
		render.Render(w, r, ErrTokenNotAllowed)
		return
	} else {
		if err := render.Render(w, r, UserResponse(user)); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
}

func UserResponse(user *model.User) *model.UserPayload {
	resp := &model.UserPayload{User: user}
	return resp
}

func requestLogin(w http.ResponseWriter, r *http.Request) {
	loginCredentials := &model.UserLoginRequest{}
	// Check Missing fields
	if err := render.Bind(r, loginCredentials); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if user, err := model.DbUpdateUserAccessToken(loginCredentials.User); err != nil {
		// internal error
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrInternal(err))
	} else {
		// ok
		render.Status(r, http.StatusOK)
		render.Render(w, r, NewTokenResponse(user))
		return
	}
}

func NewTokenResponse(user *model.User) *model.AccessToken {
	return user.AccessToken
}
