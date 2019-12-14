package models

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gosimple/slug"
)

// Index - JSON strucuture to manipulates indexes
type Index struct {
	Name  string
	Query string `json:"query"`
}

// IndexRequest - JSON strucuture to manipulates IndexRequest
type IndexRequest struct {
	*Index
}

// IndexResponse - JSON strucuture to manipulates IndexResponse
type IndexResponse struct {
	Name string
	Resp string `json:"response"`
}

// NewIndexResponse - the Response converted to type IndexResponse struct
func NewIndexResponse(indexName string, responseMessage string) *IndexResponse {
	resp := &IndexResponse{Name: indexName, Resp: responseMessage}
	return resp
}

// Render - pattern for managing payload encoding and decoding in Index Response Struct
func (rd *IndexResponse) Render(w http.ResponseWriter, r *http.Request) error {
	//rd.Name = slugify(rd.Name)
	text := strings.ReplaceAll(rd.Name, "_", "-")
	rd.Name = slug.Make(text)

	return nil
}

// Bind - Binds the Request to type IndexRequest to validate json data
func (a *IndexRequest) Bind(r *http.Request) error {
	if a.Index == nil {
		return errors.New("missing required Index fields")
	}
	a.Index.Name = strings.ToLower(a.Index.Name)
	return nil
}

func DBDeleteIndex(indexName string) (int, error) {

	/*var currManager handlers.Manager
	resp, errIndex := currManager.DeleteIndex(indexName)*/
	resp, errIndex := DeleteIndex(indexName)
	if errIndex != nil {
		return 1, errIndex
	}

	return resp, nil
}

// Mover para elasticsearch.go -> via interface
func DeleteIndex(indexName string) (int, error) {

	return 1, nil
}
