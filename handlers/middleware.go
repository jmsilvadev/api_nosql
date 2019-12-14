package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gosimple/slug"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type XKey string

// IndexCtx middleware is used to load an Index object from
// the URL parameters passed through as the request. In case
// the Index could not be found, we stop here and return a 404.
func IndexCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		indexName := chi.URLParam(r, "indexName")

		if indexName == "" {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		//indexName = slugify(indexName)
		indexName = strings.ReplaceAll(indexName, "_", "-")
		indexName = slug.Make(indexName)
		
		ctx := context.WithValue(r.Context(), XKey("indexName"), indexName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}