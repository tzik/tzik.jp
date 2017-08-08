package handler

import (
	"fmt"
	"net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "k\n")
}

func serveAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kk\n")
}

func init() {
	http.HandleFunc("/", serve)
	http.HandleFunc("/admin", serveAdmin)
}
