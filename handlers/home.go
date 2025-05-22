package handlers

import (
	"log/slog"
	"net/http"

	"github.com/kakarotDevs/brzdoors-goth/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	slog.Info("Serving homepage")
	return Render(w, r, home.Index())
}

func HandleClick(w http.ResponseWriter, r *http.Request) error {
	slog.Info("HTMX click received")
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte(`<h1 class="text-2xl text-green-600">You clicked the button!</h1>`))
	return err
}
