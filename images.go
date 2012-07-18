package main

import (
	"image"
  "image/color"
	"code.google.com/p/go-tour/pic"
)

type Image struct{
  width int
  height int
}

func (i *Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.width, i.height)
}
  
func (* Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
  r := 0xFF
  g := x*0xFF/i.width
  b := y*0xFF/i.height
  a := 0xFF
  
  return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func main() {
	m := Image{10, 10}
	pic.ShowImage(&m)
}