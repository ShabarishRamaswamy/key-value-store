package router

import (
	"encoding/json"
	"net/http"

	"github.com/shabarishramaswamy/key-value-store/src/deleteHandler"
	"github.com/shabarishramaswamy/key-value-store/src/getHandler"
	"github.com/shabarishramaswamy/key-value-store/src/insertHandler"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var kvPair models.KeyValuePair

			json.NewDecoder(r.Body).Decode(&kvPair)
			insertHandler.Insert(kvPair)

			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	router.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Get things.
			var kvPair models.KeyValuePair

			json.NewDecoder(r.Body).Decode(&kvPair)
			getHandler.Get(kvPair)

			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	router.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			// Delete things.
			var kvPair models.KeyValuePair

			json.NewDecoder(r.Body).Decode(&kvPair)
			deleteHandler.Delete(kvPair)

			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Not Allowed"))
		}
	})

	return router
}
