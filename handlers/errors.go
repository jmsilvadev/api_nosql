package handlers

import (
	"net/http"
	"github.com/go-chi/render"
)

// ErrResponse - json structure of Errors Responses
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// APIError - structure
type APIError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Render - pattern for managing payload encoding and decoding in Error Response Struct
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest - the Response when a Request is invalid
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

// ErrRender - the Response when the render fails
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response",
		ErrorText:      err.Error(),
	}
}

//ErrNotFound - When Status code is 404
var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
