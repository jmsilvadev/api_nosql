package models

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

// Info - JSON strucuture to manipultes Info
type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

// InfoResponse - JSON strucuture to manipultes InfoResponse
type InfoResponse struct {
	*Info
	SourceIP string `json:"request_ip"`
}

var Information = []*Info{
	{Title: "NoSQL Manager Service.", Description: "API RestFULL to interacte with Big Data & NoSQL DB."},
}

func dbGetInfo() (*Info, error) {
	for _, a := range Information {
		return a, nil
	}
	return nil, errors.New("Information not found")
}

// NewInfoListResponse - the Response converted to ArrayList type InfoResponse struct
func NewInfoListResponse(infos []*Info) []render.Renderer {
	list := []render.Renderer{}
	for _, info := range infos {
		list = append(list, NewInfoResponse(info))
	}
	return list
}

// NewInfoResponse - the Response converted to type InfoResponse struct
func NewInfoResponse(info *Info) *InfoResponse {
	resp := &InfoResponse{Info: info}
	return resp
}

// Render - pattern for managing payload encoding and decoding in Info Response Struct
func (rd *InfoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	rd.SourceIP = r.RemoteAddr
	return nil
}
