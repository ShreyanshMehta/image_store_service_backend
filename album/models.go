package album

import (
	"www.github.com/ShreyanshMehta/image_store_service_backend/album/image"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

type Album struct {
	Name       string                  `json:"album_name"`
	ImageList  map[string]*image.Image `json:"image_list"`
	ImageCount int                     `json:"image_count"`
	CreatedAt  string                  `json:"created_at"`
	ModifiedAt string                  `json:"modified_at"`
}

func (album *Album) getAllAlbumImages() map[string]interface{} {
	imageList := make([]interface{}, 0)
	for id, img := range album.ImageList {
		temp := img.GetImageDetail()
		temp["image_id"] = id
		imageList = append(imageList, temp)
	}
	return map[string]interface{}{
		"album_name": album.Name,
		"images":     imageList,
	}
}

func (album *Album) getAnImage(imageId string) (interface{}, bool) {
	var img interface{}
	img, isPresent := album.ImageList[imageId]
	return img, isPresent
}

func (album *Album) isImageNameAvailable(imageName string) bool {
	imageName = common.HashName(imageName)
	_, isPresent := album.ImageList[imageName]
	return isPresent
}

func (album *Album) createAnImage(imageName string) map[string]interface{} {
	keyName := common.HashName(imageName)
	album.ImageList[keyName] = &image.Image{
		Name:      imageName,
		AlbumName: album.Name,
		CreatedAt: common.GetCurrentTime(),
	}
	album.ImageCount++
	album.ModifiedAt = common.GetCurrentTime()
	data := album.ImageList[keyName].GetImageDetail()
	data["image_id"] = keyName
	return data
}

func (album *Album) deleteAnImage(imageId string) {
	delete(album.ImageList, imageId)
	album.ImageCount--
	album.ModifiedAt = common.GetCurrentTime()
}
