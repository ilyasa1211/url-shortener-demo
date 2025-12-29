package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	application "github.com/ilyasa1211/url-shortener-demo/internal/application"
)

type SiteHandler struct {
	s *application.SiteService
}
type Response struct {
	Data string `json:"data"`
}

func NewSiteHandler(s *application.SiteService) *SiteHandler {
	return &SiteHandler{s}
}

func (h *SiteHandler) Index(w http.ResponseWriter, r *http.Request) {
	sites := h.s.FindAll()

	if r.Context().Err() != nil {
		fmt.Println("Request done")
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sites)
}

func (h *SiteHandler) Show(w http.ResponseWriter, r *http.Request) {
	site, err := h.s.FindByAlias(r)

	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	if r.Context().Err() != nil {
		fmt.Println("Request done")
	}

	// w.Header().Set("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(site)

	// imidiate redirect
	w.Header().Set("Location", site)
	w.WriteHeader(http.StatusMovedPermanently)
}

func (h *SiteHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.s.Create(r)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	resp := Response{
		Data: "Created",
	}

	res, err := json.Marshal(resp)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}

func (h *SiteHandler) Update(w http.ResponseWriter, r *http.Request) {
	err := h.s.UpdateByAlias(r)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	resp := Response{
		Data: "Updated",
	}

	res, err := json.Marshal(resp)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}
func (h *SiteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.s.DeleteByAlias(r)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	resp := Response{
		Data: "Deleted",
	}

	res, err := json.Marshal(resp)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured %v", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}
