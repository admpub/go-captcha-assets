package thumbs

import (
	"embed"
	"image"
	"io"

	"github.com/admpub/go-captcha-assets/helper"
)

//go:embed *.jpg
var FS embed.FS

var Thumbs []image.Image

func GetThumbs() ([]image.Image, error) {
	var images []image.Image
	dirs, err := FS.ReadDir(".")
	if err != nil {
		return images, err
	}
	for _, dir := range dirs {
		f, err := FS.Open(dir.Name())
		if err != nil {
			return images, err
		}
		asset, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			return images, err
		}
		img, err := helper.DecodeByteToJpeg(asset)
		if err != nil {
			return images, err
		}
		images = append(images, img)
	}
	return images, nil
}

func init() {
	Thumbs, _ = GetThumbs()
}
