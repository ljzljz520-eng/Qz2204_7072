package api

import (
	"encoding/json"
	"net/http"
	"training41/domain"
	"training41/service"
	"training41/workflow"
)

type Handler struct{ S *service.Service }

func New(s *service.Service) *Handler { return &Handler{S: s} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.read(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var x domain.Record
	if json.NewDecoder(r.Body).Decode(&x) != nil {
		w.WriteHeader(400)
		return
	}
	if e := workflow.Intake(h.S, x); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(201)
}
func (h *Handler) read(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	x, e := workflow.Query(h.S, id)
	if e != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(x)
}
