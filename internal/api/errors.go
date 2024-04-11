package api

import (
	"net/http"

	"github.com/go-chi/render"
)

//Error definitions

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"code"`
	StatusText     string `json:"status"`
	ErrorText      string `json:"errorDetail,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

var InvalidUserPass = &ErrResponse{HTTPStatusCode: 403, StatusText: "Invalid user/pass."}

var ErrTokenNotAllowed = &ErrResponse{HTTPStatusCode: 403, StatusText: "Invalid or expired AccessToken"}

func ErrUnAuthorizedAction(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 310,
		StatusText:     "Unauthorized action",
		ErrorText:      err.Error(),
	}
}
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal error.",
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

func ErrProvider(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 350,
		StatusText:     "Provider error.",
		ErrorText:      err.Error(),
	}
}
