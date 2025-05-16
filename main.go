package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/kakarotDevs/brzdoors-goth/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := chi.NewMux()

	// Middleware should come before route declarations
	r.Use(middleware.Recoverer)

	r.Handle("/public/*", public())

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	r.Get("/", handlers.Make(handlers.HandleHome))
	r.Get("/about", handlers.Make(handlers.HandleAbout))
	r.Get("/contact", handlers.Make(handlers.HandleContact))
	r.Get("/order", handlers.Make(handlers.HandleOrder))

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":3000"
	}

	slog.Info("3, 2, 1, Let's Jam...", "url", fmt.Sprintf("http://localhost%s", listenAddr))
	log.Fatal(http.ListenAndServe(listenAddr, r))
}
