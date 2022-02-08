package album

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
	"www.github.com/ShreyanshMehta/image_store_service_backend/config"
)

var db map[string]*Album = make(map[string]*Album)

func createNewAlbumInDB(name string) (map[string]interface{}, error) {
	keyName := common.HashName(name)
	createdAt := common.GetCurrentTime()
	modifiedAt := common.GetCurrentTime()
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO image_store_service_albums (album_id, album_name, is_active, created_at,"+
		" modified_at) VALUES ('%s', '%s', %b, '%s', '%s')", keyName, name, 1, createdAt, modifiedAt)
	_, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"name":        name,
		"album_id":    keyName,
		"image_count": 0,
		"created_at":  createdAt,
		"modified_at": modifiedAt,
	}, nil
}

func fetchAlbumsFromDB() ([]interface{}, error) {
	albums := make([]interface{}, 0)
	db, err1 := config.ConnectDB()
	defer db.Close()
	if err1 != nil {
		return nil, err1
	}
	query := "SELECT t1.album_id, album_name, created_at, modified_at, IFNULL(image_count, 0) FROM " +
		"(SELECT album_id, album_name, created_at, modified_at " +
		"FROM image_store_service_albums WHERE is_active=1) t1 " +
		"LEFT JOIN " +
		"(SELECT album_id, count(*) as image_count " +
		"FROM image_store_service_images WHERE is_active=1 " +
		"GROUP BY album_id) t2 " +
		"ON t1.album_id=t2.album_id;"
	data, err2 := db.Query(query)
	if err2 != nil {
		return nil, err2
	}
	for data.Next() {
		album := Album{}
		err := data.Scan(&album.Id, &album.Name, &album.CreatedAt, &album.ModifiedAt, &album.ImageCount)
		if err == nil {
			albums = append(albums, album)
		}
	}
	return albums, nil
}

func deleteAlbumFromDB(albumId string) error {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE image_store_service_albums SET is_active=0 WHERE album_id='%s'", albumId)
	_, _ = db.Query(query)
	query = fmt.Sprintf("UPDATE image_store_service_images SET is_active=0 WHERE album_id='%s'", albumId)
	_, _ = db.Query(query)
	return nil
}

func isAlbumAvailableInDB(albumIdName string) (bool, error) {
	albumIdName = common.HashName(albumIdName)
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("SELECT COUNT(*) FROM image_store_service.image_store_service_albums WHERE "+
		"album_id='%s' and is_active=1", albumIdName)
	var isAvailable int
	_ = db.QueryRow(query).Scan(&isAvailable)
	if isAvailable > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetAnAlbum(albumId string) (*Album, error) {
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var album Album
	query := fmt.Sprintf("SELECT modified_at, created_at, album_id, album_name "+
		"FROM image_store_service_albums "+
		"WHERE is_active=1 and album_id='%s'", albumId)
	err2 := db.QueryRow(query).Scan(&album.ModifiedAt, &album.CreatedAt, &album.Id, &album.Name)
	if err2 != nil {
		return nil, nil
	} else {
		return &album, nil
	}
}

func Init(w http.ResponseWriter, _ *http.Request) {
	for i := 1; i < 12; i++ {
		album := "Album " + strconv.FormatInt(int64(i), 10)
		_, _ = createNewAlbumInDB(album)
		for j := 1; j < 7; j++ {
			img := "Image " + strconv.FormatInt(int64(j), 10)
			db[common.HashName(album)].createAnImage(img)
		}
	}
	_ = json.NewEncoder(w).Encode("Initialised dummy data successfully!!")
}
