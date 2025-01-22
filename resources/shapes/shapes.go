package shapes

import (
	"image"

	"github.com/admpub/go-captcha-assets/sourcedata/shapes"
)

func GetShapes() (map[string]image.Image, error) {
	return shapes.Shapes, nil
}
