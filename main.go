package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"www.github.com/ShreyanshMehta/image_store_service_backend/album"
	"www.github.com/ShreyanshMehta/image_store_service_backend/config"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(loggingMiddleware)
	config.DatabaseSchemaInit()
	setRoutes(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func setRoutes(r *mux.Router) {
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
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
