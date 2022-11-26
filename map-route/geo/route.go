package geo

import (
	"log"
	"math"
	"sort"

	"github.com/han-sz/map-route/util"
)

type Route struct {
	Size   uint64
	Routes []Sequence
}

func NewRoute() *Route {
	return &Route{Size: 0, Routes: make([]Sequence, 0, 100)}
}

func (r *Route) AsImageByteBuf(w, h int) []byte {
	routes := r.Routes[:]
	max, min := r.findBoundaries()
	buf := make([]byte, w*(h+1), w*(h+1))
	for _, v := range routes {
		for _, g := range v.Geopoints {
			p := normalise(&g, &max, &min)
			x := int(p.lat * float32(w))
			y := int(p.lon * float32(h))
			idx := y*w + x
			buf[idx] = 1

		}
	}
	return buf
}

func (r *Route) findBoundaries() (max, min geopoint) {
	var maxLat, minLat float32 = -math.MaxFloat32, math.MaxFloat32
	var maxLon, minLon float32 = -math.MaxFloat32, math.MaxFloat32

	routes := r.Routes[:]
	for _, v := range routes {
		maxLat = -util.MinFloat(-maxLat, -v.max.lat)
		maxLon = -util.MinFloat(-maxLon, -v.max.lon)
		minLat = util.MinFloat(minLat, v.min.lat)
		minLon = util.MinFloat(minLon, v.min.lon)
	}
	// max and min are used to normalise lat-lon pairs to cartesian co-ords
	max = geopoint{float32(maxLat), float32(maxLon)}
	min = geopoint{float32(minLat), float32(minLon)}
	return
}

func (r *Route) Sort() {
	sort.Slice(r.Routes, func(i, j int) bool {
		g1, g2 := r.Routes[i], r.Routes[j]
		// Sort by sequence blocks' min lat-lon pairs
		if g1.min.lat < g2.min.lat && g1.min.lon < g2.min.lon {
			return true
		}
		return false
	})
}

func (r *Route) PrintNormalised() {
	for _, v := range r.Routes {
		for _, g := range v.Geopoints {
			log.Println(normalise(&g, &v.max, &v.min))
		}
	}
}

func (r *Route) Print() {
	for _, v := range r.Routes {
		log.Println("Sequence", v.id, v.min, v.max)
	}
}
