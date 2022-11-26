package io

import (
	"encoding/csv"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/han-sz/map-route/geo"
)

func ReadRoute(r *geo.Route, path string) error {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("Route path specified was invalid at %s: %s\n", path, err)
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Println("Generating routes")
	reader := csv.NewReader(file)

	var seq *geo.Sequence
	var id uint32 = 0

	addSeq := func() {
		if seq != nil {
			log.Println("Added sequence", seq)
			r.Routes = append(r.Routes, *seq)
		}

	}
	makeSeq := func() {
		seq = geo.NewSequence(id)
		id++
	}

	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		record = record[1:]
		sequence, _ := strconv.Atoi(record[2])
		if sequence == 1 {
			addSeq()
			makeSeq()
		}
		point := geo.ToGeopoint(record[0], record[1])
		seq.Geopoints = append(seq.Geopoints, point)
		seq.Distance, _ = strconv.ParseFloat(record[3], 64)
		seq.Size++

		seq.UpdateLast()
	}
	addSeq()

	log.Println("Parsed", len(r.Routes), "route sequences")
	return nil
}
