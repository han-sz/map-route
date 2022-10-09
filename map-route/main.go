package main

import (
	"log"

	"github.com/han-sz/map-route/geo"
)

func init() {
	log.Println("Started map-route")
}

func main() {
	r := geo.NewRoute("./routes.csv")
	r.LoadAndParse()
	// r.PrintNormalised()
}
