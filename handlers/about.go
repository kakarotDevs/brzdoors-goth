package handlers

import (
	"log/slog"
	"net/http"

	"github.com/kakarotDevs/brzdoors-goth/views/about"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) error {
	slog.Info("Serving aboutpage")
	return Render(w, r, about.Index())
}
