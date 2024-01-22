//go:generate templ generate
//go:generate go fmt ./...

package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/originaldaemon/scromble/templates/pages"
	"github.com/originaldaemon/scromble/templates/partials"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	component := pages.HomePage()
	component.Render(r.Context(), w)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	component := pages.HomePage()
	component.Render(r.Context(), w)
}

func getServerStatus(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(time.Duration(rand.Intn(5)) * time.Second))
	var component = partials.ServerStatusMessageOK()
	if rand.Float32() < 0.5 {
		component = partials.ServerStatusMessageBad()
	}
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
	mux.HandleFunc("/server-status", getServerStatus)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":"+port, mux)
}
