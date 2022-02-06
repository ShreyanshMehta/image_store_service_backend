package image

import (
	"reflect"
	"testing"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

func TestImage_GetImageDetail(t *testing.T) {
	type fields struct {
		Name      string
		AlbumName string
		CreatedAt string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "Getting Image Details",
			fields: fields{
				Name:      "Image 1",
				AlbumName: "Album 1",
				CreatedAt: common.GetCurrentTime(),
			},
			want: map[string]interface{}{
				"image_name": "Image 1",
				"album_name": "Album 1",
				"created_at": common.GetCurrentTime(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				Name:      tt.fields.Name,
				AlbumName: tt.fields.AlbumName,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := img.GetImageDetail(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImageDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
