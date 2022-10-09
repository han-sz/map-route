package geo

import (
	"fmt"
	"math"
)

type sequence struct {
	id         uint32
	size       uint16
	kilometres float64

	max       geopoint
	min       geopoint
	geopoints []geopoint
}

func newSequence(id uint32) *sequence {
	return &sequence{
		id:         id,
		size:       0,
		kilometres: 0.0,

		max:       geopoint{-math.MaxFloat32, -math.MaxFloat32},
		min:       geopoint{math.MaxFloat32, math.MaxFloat32},
		geopoints: make([]geopoint, 0, 30),
	}
}

func (s *sequence) updateLast() {
	if len(s.geopoints) == 0 {
		return
	}
	last := s.geopoints[len(s.geopoints)-1]
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

func (s *sequence) String() string {
	return fmt.Sprintf("sequence[%d]{min: %v, max: %v, size: %d}", s.id, s.min, s.max, s.size)
}
