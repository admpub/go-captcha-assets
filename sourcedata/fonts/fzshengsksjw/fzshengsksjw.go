package fzshengsksjw

import (
	_ "embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

//go:embed fzshengsksjw_cu.ttf
var asset []byte

func GetFont() (font *truetype.Font, err error) {
	font, err = freetype.ParseFont(asset)
	if err != nil {
		return nil, err
	}

	return font, nil
}
