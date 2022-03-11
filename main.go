package main

import (
	"fmt"
	"image"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/fogleman/gg"
)

func main() {
	for i := 1; i < 10; i++ {
		DrawPNG(fmt.Sprintf("Pallet-%d", i))
	}
}

func DrawPNG(text string) {
	const W = 1600
	const H = 1000
	const P = 150
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", P); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(text, W/2, P, 0.5, 0, 0, 1.5, gg.AlignCenter)
	dc.DrawImage(GenerageBarCode(text), 300, 400)
	dc.SavePNG(fmt.Sprintf("out/%s.png", text))
}

func GenerageBarCode(text string) image.Image {
	barCode, _ := code128.Encode(text)
	scaledBarCode, _ := barcode.Scale(barCode.(barcode.Barcode), 1000, 300)
	return scaledBarCode
}