// NOTE: This sample intentionally omits error handling best practices, validation,
// and dependency wiring so candidates can discuss and propose improvements.

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andynuge/api-samples/badge-management/badgestore"
)

// handler holds dependencies for HTTP handlers.
type handler struct {
	badgeStore badgestore.BadgeStore
}

// registerBadgeRequest represents the expected JSON body for badge registration.
type registerBadgeRequest struct {
	SerialNumber string `json:"serialNumber"`
	Version      string `json:"version"`
}

// RegisterBadge handles the registration of a new badge.
func (h *handler) RegisterBadge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}

	var req registerBadgeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", 400)
		return
	}

	badge, err := h.badgeStore.Create(req.SerialNumber, req.Version)
	if err != nil {
		http.Error(w, "failed to create badge", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(badge)
}

// Entrypoint
func main() {
	// Initialize multiplexer for routing HTTP requests
	mux := http.NewServeMux()

	h := &handler{
		badgeStore: badgestore.New(),
	}

	mux.HandleFunc("/badges", h.RegisterBadge)

	log.Println("listening on :8080")

	http.ListenAndServe(":8080", mux)
}
