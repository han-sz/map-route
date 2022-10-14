package plot

import (
	"image"
	"image/color"
)

type Frame struct {
	Width  int
	Height int

	Buf []byte
}

func NewFrame(width, height int) *Frame {
	return &Frame{
		Buf:    make([]byte, width*height),
		Width:  width,
		Height: height,
	}
}

func (f *Frame) ConvertBufToImage() *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(f.Width), int(f.Height)}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	black := color.RGBA{0xff, 0xff, 0xff, 0xff}

	w := int(f.Width)
	h := int(f.Height)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			idx := y*w + x
			if f.Buf[idx]&0x1 != 0 {
				img.Set(x, y, black)
			}
		}
	}
	return img
}
