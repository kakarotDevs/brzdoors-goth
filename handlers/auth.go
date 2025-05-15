package handlers

import (
	"fmt"
	"net/http"

	"github.com/kakarotDevs/brzdoors-goth/views/auth"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin@example.com" && password == "secret" {
			// Could set a session here
			fmt.Fprint(w, "<p>Welcome, admin!</p>")
		} else {
			fmt.Fprint(w, "<p style='color: red;'>Invalid login</p>")
		}
		return nil
	}

	// GET: render the login form
	return Render(w, r, auth.Login())
}
