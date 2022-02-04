package image

type Image struct {
	Name      string `json:"image_name"`
	CreatedAt string `json:"created_at"`
}

func (img *Image) GetImageDetail() map[string]interface{} {
	result := map[string]interface{}{
		"image_name": img.Name,
		"created_at": img.CreatedAt,
	}
	return result
}
