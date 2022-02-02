package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"www.github.com/ShreyanshMehta/image_store_service_backend/album"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	setRoutes(r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

func setRoutes(r *mux.Router) {
	r.HandleFunc("/heath", healthCheck).Methods("GET")
	r.HandleFunc("/init", album.Init).Methods("GET")
	album.HandleAlbumRequests(r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]bool{"status": true})
	if err != nil {
		log.Fatal("Something went wrong.")
	}
	return
}
