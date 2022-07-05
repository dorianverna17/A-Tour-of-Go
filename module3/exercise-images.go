/*
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
*/

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	pixels [][]color.Color
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i.pixels[0]), len(i.pixels))
}

func (i Image) At(x, y int) color.Color {
	return i.pixels[y][x]
}

func Pic(dx, dy int) [][]color.Color {
	sl := make([][]color.Color, dy)
	for i := range sl {
		sl[i] = make([]color.Color, dx)
		for j := 0; j < dx; j++ {
			sl[i][j] = color.RGBA{uint8((i+j)/2 + i ^ j), uint8((i+j)/2 + i ^ j), 255, 255}
		}
	}
	return sl
}

func main() {
	m := Image{Pic(50, 100)}
	pic.ShowImage(m)
}
