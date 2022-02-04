package album

import (
	"www.github.com/ShreyanshMehta/image_store_service_backend/album/image"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

type Album struct {
	Name       string                  `json:"album_name"`
	ImageList  map[string]*image.Image `json:"image_list"`
	CreatedAt  string                  `json:"created_at"`
	ModifiedAt string                  `json:"modified_at"`
}

func (album *Album) getAlbumImages() []interface{} {
	imageList := make([]interface{}, 0)
	for id, img := range album.ImageList {
		temp := img.GetImageDetail()
		temp["image_id"] = id
		imageList = append(imageList, temp)
	}
	return imageList
}

func (album *Album) getAnImage(imageId string) (interface{}, bool) {
	var img interface{}
	img, isPresent := album.ImageList[imageId]
	return img, isPresent
}

func (album *Album) isImageNameAvailable(imageName string) bool {
	_, isPresent := album.ImageList[imageName]
	return isPresent
}

func (album *Album) createAnImage(imageName string) {
	keyName := common.HashName(imageName)
	album.ImageList[keyName] = &image.Image{
		Name:      imageName,
		CreatedAt: common.GetCurrentTime(),
	}
	album.ModifiedAt = common.GetCurrentTime()
}

func (album *Album) deleteAnImage(imageName string) {
	keyName := common.HashName(imageName)
	delete(album.ImageList, keyName)
	album.ModifiedAt = common.GetCurrentTime()
}
