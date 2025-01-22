package images

import (
	"image"

	"github.com/admpub/go-captcha-assets/sourcedata/images"
)

func GetImages() ([]image.Image, error) {
	images := images.Images
	return images, nil
}
