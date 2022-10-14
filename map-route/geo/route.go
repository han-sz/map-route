package geo

import (
	"encoding/csv"
	"errors"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"sort"
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

func (r *Route) Sort() {
	sort.Slice(r.routes, func(i, j int) bool {
		g1, g2 := r.routes[i], r.routes[j]
		// Sort by sequence blocks' min lat-lon pairs
		if g1.min.lat < g2.min.lat && g1.min.lon < g2.min.lon {
			return true
		}
		return false
	})
}

func (r *Route) LoadAndParse() (error, bool) {
	if _, err := os.Stat(r.path); errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("Route path specified was invalid at %s: %s\n", r.path, err)
	}
	file, err := os.Open(r.path)
	if err != nil {
		return err, false
	}
	defer file.Close()

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
func (r *Route) AsImageByteBuf(w, h int) []byte {
	var maxLat, minLat float32 = -math.MaxFloat32, math.MaxFloat32
	var maxLon, minLon float32 = -math.MaxFloat32, math.MaxFloat32

	routes := r.routes[:]
	for _, v := range routes {
		if v.max.lat > maxLat {
			maxLat = v.max.lat
		}
		if v.min.lat < minLat {
			minLat = v.min.lat
		}
		if v.max.lon > maxLon {
			maxLon = v.max.lon
		}
		if v.min.lon < minLon {
			minLon = v.min.lon
		}
	}
	// max and min are used to normalise lat-lon pairs to cartesian co-ords
	max := geopoint{float32(maxLat), float32(maxLon)}
	min := geopoint{float32(minLat), float32(minLon)}
	buf := make([]byte, w*(h+1), w*(h+1))
	for _, v := range routes {
		for _, g := range v.geopoints {
			p := normalise(&g, &max, &min)
			x := int(p.lat * float32(w))
			y := int(p.lon * float32(h))
			idx := y*w + x
			buf[idx] = 1

		}
	}
	return buf
}

func (r *Route) Print() {
	for _, v := range r.routes {
		log.Println("Sequence", v.id, v.min, v.max)
	}
}
