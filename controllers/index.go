package controllers

import (
	"log"
	"net/http"

	"github.com/go-chi/render"

	"api-nosql/handlers"
	"api-nosql/models"
)

//GetIndex - Shows Index Data
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexName := r.Context().Value(handlers.XKey("indexName")).(string)

	log.Println(indexName)
	return
}

//DeleteIndex - Delete Index Data
func DeleteIndex(w http.ResponseWriter, r *http.Request) {

	indexName := r.Context().Value(handlers.XKey("indexName")).(string)
	data := &models.IndexRequest{}
	log.Println(data)

	newMessage := "Index Deleted Successfully"

	resp, err := models.DBDeleteIndex(indexName)
	if err != nil {
		render.Render(w, r, handlers.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)

	if resp == 1 {
		newMessage = "Fail! Index Not Deleted."
	}

	render.Render(w, r, models.NewIndexResponse(indexName, newMessage))
}

//CreateIndex - Create Index Data
/*func CreateIndex(w http.ResponseWriter, r *http.Request) {

}*/
