package geo

import (
	"encoding/csv"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"
)

type Route struct {
	path   string
	size   uint64
	routes []sequence
}

type RouteInterface interface {
	Load(path string) (error, bool)
}

func NewRoute(path string) *Route {
	return &Route{path: path, size: 0, routes: make([]sequence, 0, 100)}
}

func (r *Route) LoadAndParse() (error, bool) {
	if _, err := os.Stat(r.path); errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("Route path specified was invalid at %s: %s\n", r.path, err)
	}
	file, err := os.Open(r.path)
	if err != nil {
		return err, false
	}

	log.Println("Generating routes")
	reader := csv.NewReader(file)

	var seq *sequence
	var id uint32 = 0
	addSeq := func() {
		if seq != nil {
			log.Println("Added sequence", seq)
			r.routes = append(r.routes, *seq)
		}

	}
	makeSeq := func() {
		seq = newSequence(id)
		id++
	}
	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		record = record[1:]
		sequence, _ := strconv.Atoi(record[2])
		if sequence == 1 {
			addSeq()
			makeSeq()
		}
		point := toGeopoint(record[0], record[1])
		seq.geopoints = append(seq.geopoints, point)
		seq.kilometres, _ = strconv.ParseFloat(record[3], 64)
		seq.size++

		seq.updateLast()
	}
	addSeq()
	log.Println("Parsed", len(r.routes), "route sequences")
	return nil, true
}

func (r *Route) PrintNormalised() {
	for _, v := range r.routes {
		for _, g := range v.geopoints {
			log.Println(normalise(&g, &v.max, &v.min))
		}
	}
}
