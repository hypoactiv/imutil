package imutil

import "image"

type ColorChannel uint8

const (
	RED = iota
	GREEN
	BLUE
	ALPHA
)

func Separate(in *image.RGBA, offset ColorChannel) (out *image.Gray) {
	out = image.NewGray(in.Bounds())
	for x := in.Bounds().Min.X; x < in.Bounds().Max.X; x++ {
		for y := in.Bounds().Min.Y; y < in.Bounds().Max.Y; y++ {
			out.Pix[out.PixOffset(x, y)] = in.Pix[in.PixOffset(x, y)+int(offset)]
		}
	}
	return
}
