package io

import (
	"image/png"
	"os"

	"github.com/han-sz/map-route/plot"
)

func WriteFrameImage(f *plot.Frame, filePath string) error {
	image := f.ConvertBufToImage()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = png.Encode(file, image)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
