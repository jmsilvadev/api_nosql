package controllers

import (
	"net/http"

	"github.com/go-chi/render"

	"api-nosql/models"
	"api-nosql/handlers"
)

//HealthInfo - Shows if this API is UP and Running
func HealthInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You know, to interact with the NoSQL api. :)"))
}

//ShowInfo - Shows information about this API
func ShowInfo(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, models.NewInfoListResponse(models.Information)); err != nil {
		render.Render(w, r, handlers.ErrRender(err))
		return
	}
}
