package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", Home)

	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
