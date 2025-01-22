package shapes

import (
	"embed"
	"image"
	"io"
	"strings"

	"github.com/admpub/go-captcha-assets/helper"
)

//go:embed *.png
var FS embed.FS

func GetShapes() (map[string]image.Image, error) {
	images := map[string]image.Image{}
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
		img, err := helper.DecodeByteToPng(asset)
		if err != nil {
			return images, err
		}
		key := strings.SplitN(dir.Name(), ".", 2)[0]
		images[key] = img
	}
	return images, err
}
