package tiles

import (
	"embed"
	"image"
	"io"

	"github.com/admpub/go-captcha-assets/helper"
)

//go:embed tile-1.png tile-2.png tile-3.png tile-4.png
var FSTile embed.FS

//go:embed tile-mask-*.png
var FSTileMask embed.FS

//go:embed tile-shadow-*.png
var FSTileShadow embed.FS

type GraphImage struct {
	OverlayImage image.Image
	ShadowImage  image.Image
	MaskImage    image.Image
}

func GetTiles() ([]*GraphImage, error) {
	var graphs []*GraphImage
	dirs, err := FSTile.ReadDir(".")
	if err != nil {
		return graphs, err
	}
	for _, dir := range dirs {
		name := dir.Name()
		id := name[len("tile-") : len(name)-len(".png")]
		f, err := FSTile.Open(dir.Name())
		if err != nil {
			return graphs, err
		}
		asset, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			return graphs, err
		}
		tileImg, err := helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		f, err = FSTileShadow.Open(`tile-shadow-` + id + `.png`)
		if err != nil {
			return graphs, err
		}
		asset, err = io.ReadAll(f)
		if err != nil {
			return graphs, err
		}
		tileShadowImg, err := helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		f, err = FSTileMask.Open(`tile-mask-` + id + `.png`)
		if err != nil {
			return graphs, err
		}
		asset, err = io.ReadAll(f)
		if err != nil {
			return graphs, err
		}
		tileMaskImg, err := helper.DecodeByteToPng(asset)
		if err != nil {
			return graphs, err
		}

		graphs = append(graphs, &GraphImage{
			OverlayImage: tileImg,
			ShadowImage:  tileShadowImg,
			MaskImage:    tileMaskImg,
		})
	}
	return graphs, nil
}
