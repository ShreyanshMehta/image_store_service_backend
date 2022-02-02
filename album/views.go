package album

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"www.github.com/ShreyanshMehta/image_store_service/common"
)

func HandleAlbumRequests(r *mux.Router) {
	s := r.PathPrefix("/album").Subrouter()
	s.HandleFunc("/fetch", getAlbums).Methods("GET")
	s.HandleFunc("/create", createAlbum).Methods("POST")
	s.HandleFunc("/delete", deleteAlbum).Methods("POST")

	t := s.PathPrefix("/{album_id}").Subrouter()
	t.HandleFunc("/images", fetchImagesFromAlbum).Methods("GET")
	t.HandleFunc("/add", addImageInAlbum).Methods("POST")
	t.HandleFunc("/delete", deleteImageInAlbum).Methods("POST")
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Request body cannot be empty"))
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["album_name"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Body Parameter 'album_name' was not found"))
		return
	}
	albumName := payload["album_name"].(string)
	if isAlbumNameAvailableInDB(albumName) {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Album name '" + albumName + "' already exist"))
		return
	}
	createNewAlbumInDB(albumName)
	_ = json.NewEncoder(w).Encode(common.SuccessMsg("Album created successfully"))
	return
}

func getAlbums(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(fetchAlbumsFromDB())
	return
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Request body cannot be empty"))
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["album_id"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Body parameter 'album_id' was not found"))
		return
	}
	albumId := payload["album_id"].(string)
	if !isAlbumNameAvailableInDB(albumId) {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Album name '" + albumId + "' does not exist"))
		return
	}
	deleteAlbumFromDB(albumId)
	_ = json.NewEncoder(w).Encode(common.SuccessMsg("Album deleted successfully"))
	return
}

/*********************************************************************************/
/********************************Image Related APIs*******************************/
/*********************************************************************************/

func addImageInAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, isAlbumAvailable := db[albumId]
	if !isAlbumAvailable {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("'" + albumId + "' is not available"))
		return
	}
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Request body cannot be empty"))
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["image_name"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Body parameter 'image_name' was not found"))
		return
	}
	imageName := payload["image_name"].(string)
	if album.isImageNameAvailable(imageName) {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Image name '" + imageName + "' already exists in album"))
		return
	}
	album.createAnImage(imageName)
	_ = json.NewEncoder(w).Encode(common.SuccessMsg("Image was added successfully to album'" + album.Name + "'"))
	return
}

func deleteImageInAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, isAlbumAvailable := db[albumId]
	if !isAlbumAvailable {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("'" + albumId + "' is not available"))
		return
	}
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Request body cannot be empty"))
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["image_name"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Body parameter 'image_name' was not found"))
		return
	}
	imageName := payload["image_name"].(string)
	if !album.isImageNameAvailable(imageName) {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("Image name '" + imageName + "' does not exist in album"))
		return
	}
	album.deleteAnImage(imageName)
	_ = json.NewEncoder(w).Encode(common.SuccessMsg("Image was deleted successfully from album'" + album.Name + "'"))
	return
}

func fetchImagesFromAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, isAlbumAvailable := db[albumId]
	if !isAlbumAvailable {
		_ = json.NewEncoder(w).Encode(common.ErrorMsg("'" + albumId + "' was not available"))
		return
	}
	images := album.getAlbumImages()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"images":     images,
		"status":     true,
		"tot_images": len(images),
	})
	return
}
