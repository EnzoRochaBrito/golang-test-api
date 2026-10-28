package note

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, h Handler) {
	mux.HandleFunc("POST /note", h.CreateNote)
	mux.HandleFunc("GET /note/{id}", h.GetNote)
	mux.HandleFunc("DELETE /note/{id}", h.DeleteNote)
}
