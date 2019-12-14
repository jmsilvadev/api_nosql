package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"api-nosql/controllers"
	"api-nosql/handlers"
)

var Version = "v1"
var BaseDir = "/" + Version

func ApiRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get(BaseDir+"/", controllers.HealthInfo)

	r.Route(BaseDir+"/info", func(r chi.Router) {
		r.Get("/", controllers.ShowInfo)
	})

	r.Route(BaseDir+"/license", func(r chi.Router) {
		r.Get("/", controllers.ShowLicense)
	})

	r.Route(BaseDir+"/{indexName}", func(r chi.Router) {
		r.Use(handlers.IndexCtx)
		r.Get("/", controllers.GetIndex)
		r.Delete("/", controllers.DeleteIndex)
		//r.Put("/", controllers.CreateIndex)
	})

	return r
}
