package thumbs

import (
	"image"

	"github.com/admpub/go-captcha-assets/sourcedata/thumbs"
)

func GetThumbs() ([]image.Image, error) {
	return thumbs.GetThumbs()
}
