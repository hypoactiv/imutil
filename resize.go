package imutil

import (
	"image"
	"image/color"
)

// resizes src image to have size dest
type Resize struct {
	Src  image.Image
	Dest image.Rectangle
}

func (r *Resize) ColorModel() color.Model {
	return r.Src.ColorModel()
}

func (r *Resize) Bounds() image.Rectangle {
	return r.Dest
}

func (r *Resize) At(x int, y int) color.Color {
	srcw := r.Src.Bounds().Dx()
	srch := r.Src.Bounds().Dy()
	dstw := r.Dest.Dx()
	dsth := r.Dest.Dy()
	return r.Src.At(x*srcw/dstw, y*srch/dsth)
}
