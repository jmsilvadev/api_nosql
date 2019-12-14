package controllers

import (
	"net/http"

	"github.com/go-chi/render"

	"api-nosql/handlers"
	"api-nosql/models"
)

//ShowLicense - Shows information about the License of this API
func ShowLicense(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, models.NewLicenseListResponse(models.Licenseinfo)); err != nil {
		render.Render(w, r, handlers.ErrRender(err))
		return
	}
}
