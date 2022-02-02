package album

import (
	"encoding/json"
	"net/http"
	"time"
	"www.github.com/ShreyanshMehta/image_store_service/album/image"
	"www.github.com/ShreyanshMehta/image_store_service/common"
)

var db map[string]*Album = make(map[string]*Album)

func createNewAlbumInDB(name string) {
	keyName := common.HashName(name)
	db[keyName] = &Album{
		Name:       name,
		ImageList:  map[string]*image.Image{},
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}

func fetchAlbumsFromDB() interface{} {
	albums := make([]interface{}, 0)
	for key, value := range db {
		albums = append(albums, map[string]interface{}{
			"name":        value.Name,
			"album_id":    key,
			"created_at":  value.CreatedAt,
			"modified_at": value.ModifiedAt,
		})
	}
	return albums
}

func deleteAlbumFromDB(id string) {
	delete(db, id)
}

func isAlbumNameAvailableInDB(name string) bool {
	name = common.HashName(name)
	_, isPresent := db[name]
	return isPresent
}

func Init(w http.ResponseWriter, _ *http.Request) {
	s := []string{"Album 1", "Album 2", "Album 3"}
	for _, a := range s {
		createNewAlbumInDB(a)
	}
	db["album1"].createAnImage("image1")
	db["album1"].createAnImage("image2")
	db["album1"].createAnImage("image3")
	db["album2"].createAnImage("image1")
	db["album3"].createAnImage("image1")
	_ = json.NewEncoder(w).Encode("Initialised dummy data successfully!!")
}
