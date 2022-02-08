package album

import (
	"reflect"
	"testing"
	"www.github.com/ShreyanshMehta/image_store_service_backend/album/image"
	"www.github.com/ShreyanshMehta/image_store_service_backend/common"
)

func TestAlbum_createAnImage(t *testing.T) {
	type fields struct {
		Name       string
		ImageList  map[string]*image.Image
		ImageCount int
		CreatedAt  string
		ModifiedAt string
	}
	type args struct {
		imageName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "Creating an image",
			fields: fields{
				Name:       "Album 1",
				ImageList:  map[string]*image.Image{},
				ImageCount: 0,
				CreatedAt:  common.GetCurrentTime(),
				ModifiedAt: common.GetCurrentTime(),
			},
			args: args{"Image 1"},
			want: map[string]interface{}{
				"album_name": "Album 1",
				"created_at": common.GetCurrentTime(),
				"image_id":   "image1",
				"image_name": "Image 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := &Album{
				Name:       tt.fields.Name,
				ImageList:  tt.fields.ImageList,
				ImageCount: tt.fields.ImageCount,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}
			if got := album.createAnImage(tt.args.imageName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createAnImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_deleteAnImage(t *testing.T) {
	album := Album{
		Name:       "Album 1",
		ImageList:  map[string]*image.Image{},
		ImageCount: 0,
		CreatedAt:  common.GetCurrentTime(),
		ModifiedAt: common.GetCurrentTime(),
	}
	_ = album.createAnImage("Image 1")
	type fields struct {
		Name       string
		ImageList  map[string]*image.Image
		ImageCount int
		CreatedAt  string
		ModifiedAt string
	}
	type args struct {
		imageId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Deleting an image",
			fields: fields{
				Name:       album.Name,
				ImageList:  album.ImageList,
				ImageCount: album.ImageCount,
				CreatedAt:  album.CreatedAt,
				ModifiedAt: album.ModifiedAt,
			},
			args: args{
				"image1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := &Album{
				Name:       tt.fields.Name,
				ImageList:  tt.fields.ImageList,
				ImageCount: tt.fields.ImageCount,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}
			album.deleteAnImage(tt.args.imageId)
		})
	}
}

func TestAlbum_getAllAlbumImages(t *testing.T) {
	type fields struct {
		Name       string
		ImageList  map[string]*image.Image
		ImageCount int
		CreatedAt  string
		ModifiedAt string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := &Album{
				Name:       tt.fields.Name,
				ImageList:  tt.fields.ImageList,
				ImageCount: tt.fields.ImageCount,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}
			if got := album.getAllAlbumImages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllAlbumImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlbum_getAnImage(t *testing.T) {
	album := Album{
		Name:       "Album 1",
		ImageList:  map[string]*image.Image{},
		ImageCount: 0,
		CreatedAt:  common.GetCurrentTime(),
		ModifiedAt: common.GetCurrentTime(),
	}
	_ = album.createAnImage("Image 1")
	type fields struct {
		Name       string
		ImageList  map[string]*image.Image
		ImageCount int
		CreatedAt  string
		ModifiedAt string
	}
	type args struct {
		imageId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name: "Getting details for an available image",
			fields: fields{
				Name:       album.Name,
				ImageList:  album.ImageList,
				ImageCount: album.ImageCount,
				CreatedAt:  album.CreatedAt,
				ModifiedAt: album.ModifiedAt,
			},
			args: args{"Image 1"},
			want: &image.Image{
				"Image 1",
				"Album 1",
				"02-07-2022 18:27:51 Monday",
			},
			want1: true,
		},
		{
			name: "Getting details for an unavailable image",
			fields: fields{
				Name:       album.Name,
				ImageList:  album.ImageList,
				ImageCount: album.ImageCount,
				CreatedAt:  album.CreatedAt,
				ModifiedAt: album.ModifiedAt,
			},
			args:  args{"Image 2"},
			want:  album.ImageList["image2"],
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := &Album{
				Name:       tt.fields.Name,
				ImageList:  tt.fields.ImageList,
				ImageCount: tt.fields.ImageCount,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}
			got, got1 := album.getAnImage(tt.args.imageId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAnImage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getAnImage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAlbum_isImageNameAvailable(t *testing.T) {
	album := Album{
		Name:       "Album 1",
		ImageList:  map[string]*image.Image{},
		ImageCount: 0,
		CreatedAt:  common.GetCurrentTime(),
		ModifiedAt: common.GetCurrentTime(),
	}
	_ = album.createAnImage("Image 1")
	type fields struct {
		Name       string
		ImageList  map[string]*image.Image
		ImageCount int
		CreatedAt  string
		ModifiedAt string
	}
	type args struct {
		imageName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "If image is available",
			fields: fields{
				Name:       album.Name,
				ImageList:  album.ImageList,
				ImageCount: album.ImageCount,
				CreatedAt:  album.CreatedAt,
				ModifiedAt: album.ModifiedAt,
			},
			args: args{
				"image1",
			},
			want: true,
		},
		{
			name: "If image is available",
			fields: fields{
				Name:       album.Name,
				ImageList:  album.ImageList,
				ImageCount: album.ImageCount,
				CreatedAt:  album.CreatedAt,
				ModifiedAt: album.ModifiedAt,
			},
			args: args{
				"image22",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := &Album{
				Name:       tt.fields.Name,
				ImageList:  tt.fields.ImageList,
				ImageCount: tt.fields.ImageCount,
				CreatedAt:  tt.fields.CreatedAt,
				ModifiedAt: tt.fields.ModifiedAt,
			}
			if got := album.isImageNameAvailable(tt.args.imageName); got != tt.want {
				t.Errorf("isImageNameAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
