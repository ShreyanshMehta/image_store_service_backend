package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"www.github.com/ShreyanshMehta/image_store_service_backend/album"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(loggingMiddleware)
	setRoutes(r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

func setRoutes(r *mux.Router) {
	r.HandleFunc("/heath", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/init", album.Init).Methods("GET")
	album.HandleAlbumRequests(r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
