package router

import "net/http"

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Insert things.
			w.Write([]byte("NYI"))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	router.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Get things.
			w.Write([]byte("NYI"))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	router.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			// Delete things.
			w.Write([]byte("NYI"))
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	return router
}
