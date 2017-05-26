package imutil

import (
	"image"
	"image/draw"
)

// Return a grayscale copy of i
func ToGrayscale(i image.Image) (out *image.Gray) {
	r := i.Bounds()
	gray := image.NewGray(r)
	draw.Draw(gray, r, i, image.ZP, draw.Src)
	return gray
}

// Return an RGBA copy of i
func ToRGBA(i image.Image) (out *image.RGBA) {
	r := i.Bounds()
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, r, i, image.ZP, draw.Src)
	return rgba
}

// Combine the grayscale images into RGB (opaque alpha channel)
func CombineRGB(r, g, b *image.Gray) (out *image.RGBA) {
	inputs := []*image.Gray{r, g, b}
	for i := 1; i < len(inputs); i++ {
		if inputs[i].Bounds() != inputs[0].Bounds() {
			panic("image bounds do not match")
		}
	}
	out = image.NewRGBA(inputs[0].Bounds())
	for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
		for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
			offset1 := out.PixOffset(x, y)
			offset2 := r.PixOffset(x, y)
			out.Pix[offset1+RED] = r.Pix[offset2]
			out.Pix[offset1+GREEN] = g.Pix[offset2]
			out.Pix[offset1+BLUE] = b.Pix[offset2]
			out.Pix[offset1+ALPHA] = 255
		}
	}
	return
}

// Combine the grayscale images into RGBA
func CombineRGBA(r, g, b, a *image.Gray) (out *image.RGBA) {
	inputs := []*image.Gray{r, g, b, a}
	for i := 1; i < len(inputs); i++ {
		if inputs[i].Bounds() != inputs[0].Bounds() {
			panic("image bounds do not match")
		}
	}
	out = image.NewRGBA(inputs[0].Bounds())
	for x := out.Bounds().Min.X; x < out.Bounds().Max.X; x++ {
		for y := out.Bounds().Min.Y; y < out.Bounds().Max.Y; y++ {
			offset1 := out.PixOffset(x, y)
			offset2 := r.PixOffset(x, y)
			out.Pix[offset1+RED] = r.Pix[offset2]
			out.Pix[offset1+GREEN] = g.Pix[offset2]
			out.Pix[offset1+BLUE] = b.Pix[offset2]
			out.Pix[offset1+ALPHA] = a.Pix[offset2]
		}
	}
	return
}
