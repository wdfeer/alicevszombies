package ui

import "github.com/gen2brain/raylib-go/raygui"

type FontSize = uint8

const (
	TextSmall FontSize = iota
	TextBig
)

var defaultFontSizeMap = map[FontSize]float32{
	TextSmall: 32,
	TextBig:   80,
}

var fontSizeMap = make(map[FontSize]float32)

func SetFontSize(fontSize FontSize) {
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, int64(fontSizeMap[fontSize]))
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SPACING, int64(fontSizeMap[fontSize]/10))
}

func updateFontSizes() {
	for k, v := range defaultFontSizeMap {
		fontSizeMap[k] = v * UIScale
	}
}
