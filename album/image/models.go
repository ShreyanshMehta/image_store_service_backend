package image

type Image struct {
	Id        string `json:"image_id"`
	Name      string `json:"image_name"`
	AlbumName string `json:"album_name"`
	CreatedAt string `json:"created_at"`
}
