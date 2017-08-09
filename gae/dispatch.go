package dispatcher

import (
	"fmt"
	"net/http"

	"github.com/tzik/tzik.jp/handler"
)

func serveAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "kk\n")
}

func init() {
	http.Handle("/", &handler.Handler{})
	http.HandleFunc("/admin/", serveAdmin)
}
