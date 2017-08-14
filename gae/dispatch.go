package dispatcher

import (
	"net/http"

	"github.com/tzik/tzik.jp/handler"
)

// func serveAdmin(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/admin/autocert" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	cert.Request()
// }

func init() {
	http.Handle("/", &handler.Handler{})
	// http.HandleFunc("/.well-known/acme-challenge/", cert.HandleHTTP01Challenge)
	// http.HandleFunc("/admin/cert", cert.RenewCertificate)
	// http.HandleFunc("/admin/", serveAdmin)
}
