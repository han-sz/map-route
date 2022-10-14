package main

import (
	"log"

	"github.com/han-sz/map-route/geo"
	"github.com/han-sz/map-route/io"
	"github.com/han-sz/map-route/plot"
)

const (
	FRAME_WIDTH  int = 1920
	FRAME_HEIGHT int = 1080

	OUTPUT_FILE string = "generated-plot.png"
)

func init() {
	log.Println("Started map-route")
}

func main() {
	// TODO: pass in as input args
	r := geo.NewRoute("./data/routes.csv")
	r.LoadAndParse()
	r.Sort()
	r.Print()

	frame := plot.NewFrame(FRAME_WIDTH, FRAME_HEIGHT)
	buf := r.AsImageByteBuf(FRAME_WIDTH, FRAME_HEIGHT)
	frame.Buf = buf

	if err := io.WriteFrameImage(frame, OUTPUT_FILE); err != nil {
		log.Panic("Failed to generated map plot:", err)
	} else {
		log.Println("Generated map plot:", OUTPUT_FILE)
	}
}
