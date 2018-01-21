package geo

import "math"

// Location represents a location in earth specified by the latitude and
// longitude.
type Location struct {
	Lat  float64
	Long float64
}

const (
	// EARTHRADIUS in radius according to wikipedia is about 6,371km
	EARTHRADIUS = 6371
)

// GreatCircleDistance calculates the Haversine distance between two points in kilometers.
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (l *Location) GreatCircleDistance(destination *Location) float64 {
	dLat := (destination.Lat - l.Lat) * (math.Pi / 180.0)
	dLon := (destination.Long - l.Long) * (math.Pi / 180.0)

	lat1 := l.Lat * (math.Pi / 180.0)
	lat2 := destination.Lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTHRADIUS * c
}
