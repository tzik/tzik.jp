package dispatcher

import (
	"fmt"
	"net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	h := w.Header()
	h.Set("content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, `<!doctype html>k`)
}

func serveAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kk\n")
}

func init() {
	http.HandleFunc("/", serve)
	http.HandleFunc("/admin/", serveAdmin)
}
