package yrdzst

import (
	"github.com/admpub/go-captcha-assets/sourcedata/fonts/yrdzst"
	"github.com/golang/freetype/truetype"
)

func GetFont() (font *truetype.Font, err error) {
	return yrdzst.GetFont()
}
