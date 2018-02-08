package main

import "github.com/fogleman/gg"

func main() {
	const S = 512
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(255, 255, 255)
	if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 64); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}
