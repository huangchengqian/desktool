package ui

import (
	"desktool/src/util"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/canvas"
	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
)

var resourceImagePng = &fyne.StaticResource{
	StaticName: "translate.png",
	StaticContent: util.Convert2Byte("C:/code/go/desktool/resource/img/translate.png"),
}

var directPng = &fyne.StaticResource{
	StaticName: "replace.png",
	StaticContent: util.Convert2Byte("C:/code/go/desktool/resource/img/replace.png"),
}

func init() {
	// 资源初始化
    pic = canvas.NewImageFromResource(resourceImagePng)
    direct = canvas.NewImageFromResource(directPng)
	// 字体初始化
	fontPath, err := findfont.Find("C:/code/go/desktool/resource/font/simkai.ttf")
	if err != nil {
		panic(err)
	}

	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}
