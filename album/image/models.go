package image

type Image struct {
	Name      string `json:"image_name"`
	AlbumName string `json:"album_name"`
	CreatedAt string `json:"created_at"`
}

func (img *Image) GetImageDetail() map[string]interface{} {
	result := map[string]interface{}{
		"image_name": img.Name,
		"album_name": img.AlbumName,
		"created_at": img.CreatedAt,
	}
	return result
}
