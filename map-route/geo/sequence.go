package geo

import (
	"fmt"
	"math"
)

type Sequence struct {
	id       uint32
	Size     uint16
	Distance float64

	max       geopoint
	min       geopoint
	Geopoints []geopoint
}

func NewSequence(id uint32) *Sequence {
	return &Sequence{
		id:       id,
		Size:     0,
		Distance: 0.0,

		max:       geopoint{-math.MaxFloat32, -math.MaxFloat32},
		min:       geopoint{math.MaxFloat32, math.MaxFloat32},
		Geopoints: make([]geopoint, 0, 30),
		// TODO: run a stat / histogram to find avg. num co-ords per sequence to save on mem alloc and copy
	}
}

func (s *Sequence) UpdateLast() {
	if len(s.Geopoints) == 0 {
		return
	}
	last := s.Geopoints[len(s.Geopoints)-1]
	if last.lat < s.min.lat {
		s.min.lat = last.lat
	} else if last.lat > s.max.lat {
		s.max.lat = last.lat
	}
	if last.lon < s.min.lon {
		s.min.lon = last.lon
	} else if last.lon > s.max.lon {
		s.max.lon = last.lon
	}
}

func (s *Sequence) String() string {
	return fmt.Sprintf("sequence[%d]{min: %v, max: %v, distance: %0.0f}", s.id, s.min, s.max, s.Distance)
}
