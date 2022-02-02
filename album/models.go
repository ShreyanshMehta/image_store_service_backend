package album

import (
	"time"
	"www.github.com/ShreyanshMehta/image_store_service/album/image"
	"www.github.com/ShreyanshMehta/image_store_service/common"
)

type Album struct {
	Name       string                  `json:"album_name"`
	ImageList  map[string]*image.Image `json:"image_list"`
	CreatedAt  time.Time               `json:"created_at"`
	ModifiedAt time.Time               `json:"modified_at"`
}

func (album *Album) getAlbumImages() []interface{} {
	imageList := make([]interface{}, 0)
	for _, img := range album.ImageList {
		imageList = append(imageList, img.GetImageDetail())
	}
	return imageList
}

func (album *Album) isImageNameAvailable(imageName string) bool {
	_, isPresent := album.ImageList[imageName]
	return isPresent
}

func (album *Album) createAnImage(imageName string) {
	keyName := common.HashName(imageName)
	album.ImageList[keyName] = &image.Image{
		Name:      imageName,
		CreatedAt: time.Now(),
	}
}

func (album *Album) deleteAnImage(imageName string) {
	keyName := common.HashName(imageName)
	delete(album.ImageList, keyName)
}
