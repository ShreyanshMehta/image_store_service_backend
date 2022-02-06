package album

import (
	"www.github.com/ShreyanshMehta/image_store_service_backend/album/image"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

var db map[string]*Album = make(map[string]*Album)

func createNewAlbumInDB(name string) map[string]interface{} {
	keyName := common.HashName(name)
	db[keyName] = &Album{
		Name:       name,
		ImageList:  map[string]*image.Image{},
		CreatedAt:  common.GetCurrentTime(),
		ModifiedAt: common.GetCurrentTime(),
	}
	return map[string]interface{}{
		"name":        db[keyName].Name,
		"album_id":    keyName,
		"image_count": db[keyName].ImageCount,
		"created_at":  db[keyName].CreatedAt,
		"modified_at": db[keyName].ModifiedAt,
	}
}

func fetchAlbumsFromDB() []interface{} {
	albums := make([]interface{}, 0)
	for key, value := range db {
		albums = append(albums, map[string]interface{}{
			"name":        value.Name,
			"album_id":    key,
			"image_count": value.ImageCount,
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

//func Init(w http.ResponseWriter, _ *http.Request) {
//	for i := 1; i < 12; i++ {
//		album := "Album " + strconv.FormatInt(int64(i), 10)
//		createNewAlbumInDB(album)
//		for j := 1; j < 7; j++ {
//			img := "Image " + strconv.FormatInt(int64(j), 10)
//			db[common.HashName(album)].createAnImage(img)
//		}
//	}
//	_ = json.NewEncoder(w).Encode("Initialised dummy data successfully!!")
//}
