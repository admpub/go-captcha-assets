package fzshengsksjw

import (
	"github.com/admpub/go-captcha-assets/sourcedata/fonts/fzshengsksjw"
	"github.com/golang/freetype/truetype"
)

func GetFont() (font *truetype.Font, err error) {
	return fzshengsksjw.GetFont()
}
