package note

import (
	"api/internal/json"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var tempNote CreateNote
	err := json.Read(r, &tempNote)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdNote, err := h.service.CreateNote(r.Context(), tempNote)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 200)
	}
	json.Write(w, http.StatusCreated, createdNote)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	noteId := r.PathValue("id")
	note, err := h.service.GetNote(r.Context(), noteId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	json.Write(w, http.StatusFound, note)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteId := r.PathValue("id")
	err := h.service.DeleteNote(r.Context(), noteId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	json.Write(w, http.StatusOK, "{success: true}")
}
