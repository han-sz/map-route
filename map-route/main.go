package main

import (
	"flag"
	"log"

	"github.com/han-sz/map-route/geo"
	"github.com/han-sz/map-route/io"
	"github.com/han-sz/map-route/plot"
)

const (
	DEFAULT_INPUT_FILE   string = "data/routes.csv"
	DEFAULT_OUTPUT_FILE  string = "generated-plot.png"
	DEFAULT_FRAME_WIDTH  int    = 1920
	DEFAULT_FRAME_HEIGHT int    = 1080
)

type options struct {
	inputData  string
	outputFile string
	width      int
	height     int
}

var opt options = options{}

func init() {
	flag.StringVar(&opt.inputData, "i", DEFAULT_INPUT_FILE, "the output file path for the generated map plot")
	flag.StringVar(&opt.outputFile, "o", DEFAULT_OUTPUT_FILE, "the output file path for the generated map plot")
	flag.IntVar(&opt.width, "w", DEFAULT_FRAME_WIDTH, "the width of the map plot")
	flag.IntVar(&opt.height, "h", DEFAULT_FRAME_HEIGHT, "the height of the map plot")

	log.Println("Started map-route")
}

func main() {
	flag.Parse()

	log.Printf("Running with flags %+v", opt)

	route := geo.NewRoute()
	if err := io.ReadRoute(route, opt.inputData); err != nil {
		log.Panic("Could not read input route data:", err)
	}
	route.Sort()
	route.Print()

	frame := plot.NewFrame(opt.width, opt.height)
	buf := route.AsImageByteBuf(opt.width, opt.height)

	frame.Buf = buf

	if err := io.WriteFrameImage(frame, opt.outputFile); err != nil {
		log.Panic("Failed to generated map plot:", err)
	}

	log.Println("Generated map plot:", opt.outputFile)
}
