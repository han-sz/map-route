package geo

import (
	"strconv"

	"github.com/han-sz/map-route/util"
)

var nullGeopoint geopoint = geopoint{-1.0, -1.0}

type geopoint struct {
	lat float32
	lon float32
}

func ToGeopoint(lat, lon string) geopoint {
	latVal, errLat := strconv.ParseFloat(lat, 32)
	lonVal, errLon := strconv.ParseFloat(lon, 32)
	if errLat != nil || errLon != nil {
		return nullGeopoint
	}
	return geopoint{
		lat: float32(latVal),
		lon: float32(lonVal),
	}
}

func normalise(g, min, max *geopoint) geopoint {
	if g == nil || min == nil || max == nil {
		return nullGeopoint
	}
	return geopoint{
		lat: util.Normalise(g.lat, min.lat, max.lat),
		lon: util.Normalise(g.lon, min.lon, max.lon),
	}
}
