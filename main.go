//go:generate templ generate
//go:generate go fmt ./...

package main

import (
	"net/http"
	"os"

	"github.com/originaldaemon/scromble/templates/pages"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	component := pages.HomePage()
	component.Render(r.Context(), w)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	component := pages.HomePage()
	component.Render(r.Context(), w)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/get-info", getInfo)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":"+port, mux)
}
