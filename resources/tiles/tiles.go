package tiles

import (
	"github.com/admpub/go-captcha-assets/sourcedata/tiles"
)

func GetTiles() ([]*tiles.GraphImage, error) {
	return tiles.Tiles, nil
}
