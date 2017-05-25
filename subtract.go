package imutil

import (
	"image"
	"image/color"
)

// Subtracts two image's RGBA data
type SubtractRGBA struct {
	// Images to subtract
	I [2]*image.RGBA
	// Image offsets
	O [2]image.Point
}

func NewSubtractRGBA(i1, i2 *image.RGBA) *SubtractRGBA {
	return NewSubtractRGBAOffset(i1, i2, image.ZP, image.ZP)
}

func NewSubtractRGBAOffset(i1, i2 *image.RGBA, o1, o2 image.Point) *SubtractRGBA {
	return &SubtractRGBA{
		I: [2]*image.RGBA{i1, i2},
		O: [2]image.Point{o1, o2},
	}
}

func (s *SubtractRGBA) ColorModel() color.Model {
	return color.RGBAModel
}

func (s *SubtractRGBA) Bounds() image.Rectangle {
	return s.I[0].Bounds().Add(s.O[0]).Union(s.I[1].Bounds().Add(s.O[1]))
}

func absdiff(a, b uint8) uint8 {
	if a > b {
		return uint8(a - b)
	}
	return uint8(b - a)
}

func (s *SubtractRGBA) At(x int, y int) color.Color {
	// read pixels
	cp := s.I[0].RGBAAt(x-s.O[0].X, y-s.O[0].Y)
	cn := s.I[1].RGBAAt(x-s.O[1].X, y-s.O[1].Y)
	// subtract and return
	return color.RGBA{
		R: absdiff(cp.R, cn.R),
		G: absdiff(cp.G, cn.G),
		B: absdiff(cp.B, cn.B),
		A: 0xff, //cp.A - cn.A,
	}
}
