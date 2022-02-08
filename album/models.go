package album

import (
	"fmt"
	"www.github.com/ShreyanshMehta/image_store_service_backend/album/image"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
	"www.github.com/ShreyanshMehta/image_store_service_backend/config"
)

type Album struct {
	Name       string `json:"name"`
	Id         string `json:"album_id"`
	ImageCount int    `json:"image_count"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
}

func (album *Album) getAllAlbumImages() (map[string]interface{}, error) {
	imageList := make([]interface{}, 0)
	db, err1 := config.ConnectDB()
	defer db.Close()
	if err1 != nil {
		return nil, err1
	}
	query := fmt.Sprintf("SELECT created_at, image_id, image_name "+
		"FROM image_store_service_images "+
		"WHERE is_active=1 and album_id='%s'", album.Id)
	data, err2 := db.Query(query)
	if err2 != nil {
		return nil, err2
	}
	for data.Next() {
		img := image.Image{}
		img.AlbumName = album.Name
		err := data.Scan(&img.CreatedAt, &img.Id, &img.Name)
		if err != nil {
			return nil, err
		}
		imageList = append(imageList, img)
	}
	album.ImageCount = len(imageList)
	return map[string]interface{}{
		"album_name": album.Name,
		"images":     imageList,
	}, nil
}

func (album *Album) getAnImage(imageId string) (interface{}, error) {
	var img image.Image
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT created_at, image_id, image_name "+
		"FROM image_store_service_images "+
		"WHERE is_active=1 and album_id='%s' and image_id='%s'", album.Id, imageId)
	err = db.QueryRow(query).Scan(&img.CreatedAt, &img.Id, &img.Name)
	if err != nil {
		return nil, nil
	}
	img.AlbumName = album.Name
	return img, nil
}

func (album *Album) isImageNameAvailable(imageName string) (bool, error) {
	imageId := common.HashName(imageName)
	var img image.Image
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("SELECT created_at, image_id, image_name "+
		"FROM image_store_service_images "+
		"WHERE is_active=1 and album_id='%s' and image_id='%s'", album.Id, imageId)
	err = db.QueryRow(query).Scan(&img.CreatedAt, &img.Id, &img.Name)
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}

}

func (album *Album) createAnImage(imageName string) (map[string]interface{}, error) {
	imageId := common.HashName(imageName)
	createdAt := common.GetCurrentTime()
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO image_store_service_images (album_id, created_at, image_id, image_name, "+
		"is_active) VALUES('%s', '%s', '%s', '%s', 1)", album.Id, createdAt, imageId, imageName)
	_, err2 := db.Exec(query)
	if err2 != nil {
		return nil, err2
	}
	query = fmt.Sprintf("UPDATE image_store_service_albums SET modified_at='%s' WHERE album_id='%s'",
		createdAt, album.Id)
	_, _ = db.Exec(query)
	return map[string]interface{}{
		"album_id":   album.Id,
		"image_id":   imageId,
		"image_name": imageName,
		"created_at": createdAt,
	}, nil
}

func (album *Album) deleteAnImage(imageId string) error {
	ModifiedAt := common.GetCurrentTime()
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE image_store_service_images SET is_active=0 "+
		"WHERE album_id='%s' and image_id='%s'", album.Id, imageId)
	_, err2 := db.Exec(query)
	if err2 != nil {
		return err2
	}
	query = fmt.Sprintf("UPDATE image_store_service_albums SET modified_at='%s' WHERE album_id='%s'",
		ModifiedAt, album.Id)
	_, _ = db.Exec(query)
	return nil
}
