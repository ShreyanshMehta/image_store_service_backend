package album

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

func HandleAlbumRequests(r *mux.Router) {
	s := r.PathPrefix("/album").Subrouter()
	s.HandleFunc("", getAlbums).Methods("GET")
	s.HandleFunc("/create", createAlbum).Methods("POST")
	s.HandleFunc("/delete", deleteAlbum).Methods("POST")

	t := s.PathPrefix("/{album_id}").Subrouter()
	t.HandleFunc("/images", fetchImagesFromAlbum).Methods("GET")
	t.HandleFunc("/add", addImageInAlbum).Methods("POST")
	t.HandleFunc("/delete", deleteImageInAlbum).Methods("POST")
	t.HandleFunc("/images/{image_id}", getImageInAlbum).Methods("GET")
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Request body cannot be empty"}.Error())
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["album_name"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Body Parameter 'album_name' was not found"}.Error())
		return
	}
	albumName := payload["album_name"].(string)
	isAvailable, err := isAlbumAvailableInDB(albumName)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to complete operation"}.Error())
		return
	}
	if isAvailable {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Album name '" + albumName + "' already exist"}.Error())
		return
	}
	data, err2 := createNewAlbumInDB(albumName)
	if err2 != nil {
		log.Println(err2)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to complete the operation"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(common.Response{Message: "Album created successfully"}.Success(data))
	return
}

func getAlbums(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := fetchAlbumsFromDB()
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch album"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(common.Response{}.Success(data))
	return
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Request body cannot be empty"}.Error())
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["album_id"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Body parameter 'album_id' was not found"}.Error())
		return
	}
	albumId := payload["album_id"].(string)
	isAvailable, err := isAlbumAvailableInDB(albumId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to complete operation"}.Error())
		return
	}
	if !isAvailable {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Album '" + albumId + "' does not exist"}.Error())
		return
	}
	err = deleteAlbumFromDB(albumId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to complete operation."}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(common.Response{Message: "Album deleted successfully"}.Success(nil))
	return
}

/*********************************************************************************/
/********************************Image Related APIs*******************************/
/*********************************************************************************/

func addImageInAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, err := GetAnAlbum(albumId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	if album == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Album not found for given album id'" + albumId + "'"}.Error())
		return
	}
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Request body cannot be empty"}.Error())
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["image_name"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Body parameter 'image_name' was not found"}.Error())
		return
	}
	imageName := payload["image_name"].(string)
	isPresent, err2 := album.isImageNameAvailable(imageName)
	if err2 != nil {
		log.Println(err2)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	if isPresent {
		_ = json.NewEncoder(w).Encode(
			common.Response{Message: "Image name '" + imageName + "' already exists in album"}.Error())
		return
	}
	data, err3 := album.createAnImage(imageName)
	if err3 != nil {
		log.Println(err3)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(
		common.Response{Message: "Image was added successfully to album '" + album.Name + "'"}.Success(data))
	return
}

func deleteImageInAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, err := GetAnAlbum(albumId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	if album == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Album not found for given album id'" + albumId + "'"}.Error())
		return
	}
	if r.Body == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Request body cannot be empty"}.Error())
		return
	}
	var payload map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&payload)
	if _, isPresent := payload["image_id"]; !isPresent {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Body parameter 'image_id' was not found"}.Error())
		return
	}
	imageId := payload["image_id"].(string)
	isAvailable, err2 := album.isImageNameAvailable(imageId)
	if err2 != nil {
		log.Println(err2)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	if !isAvailable {
		_ = json.NewEncoder(w).Encode(
			common.Response{Message: "Image ID '" + imageId + "' does not exist in album"}.Error())
		return
	}
	err2 = album.deleteAnImage(imageId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(
		common.Response{Message: "Image was deleted successfully from album '" + album.Name + "'"}.Success(nil))
	return
}

func fetchImagesFromAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	album, err1 := GetAnAlbum(albumId)
	if err1 != nil {
		log.Println(err1)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	if album == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Album not found for given album id'" + albumId + "'"}.Error())
		return
	}
	data, err2 := album.getAllAlbumImages()
	if err2 != nil {
		log.Println(err1)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the images"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(common.Response{}.Success(data))
	return
}

func getImageInAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	albumId := vars["album_id"]
	imageId := vars["image_id"]
	album, err := GetAnAlbum(albumId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the image"}.Error())
		return
	}
	if album == nil {
		_ = json.NewEncoder(w).Encode(common.Response{Message: "'" + albumId + "' is not available"}.Error())
		return
	}
	img, err2 := album.getAnImage(imageId)
	if err2 != nil {
		log.Println(err2)
		_ = json.NewEncoder(w).Encode(common.Response{Message: "Unable to fetch the image"}.Error())
		return
	}
	if img == nil {
		_ = json.NewEncoder(w).Encode(
			common.Response{Message: "Image with image_id'" + imageId + "' is not available"}.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(common.Response{}.Success([]interface{}{img}))
	return
}
