package imutil

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestSeparate(t *testing.T) {
	rgba := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{color.RGBA{255, 192, 128, 64}}, image.ZP, draw.Src)
	check := func(in *image.Gray, want uint8) {
		if in.Bounds() != rgba.Bounds() {
			t.Fatal("color channel has wrong size")
		}
		for _, pixel := range in.Pix {
			if pixel != want {
				t.Fatal("color channel has wrong value")
			}
		}
	}
	r := Separate(rgba, RED)
	g := Separate(rgba, GREEN)
	b := Separate(rgba, BLUE)
	a := Separate(rgba, ALPHA)
	check(r, 255)
	check(g, 192)
	check(b, 128)
	check(a, 64)
}
