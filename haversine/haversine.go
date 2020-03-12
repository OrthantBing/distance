package haversine

import "math"

const (
	// Radius of Earth
	miEarthRadius = 3958 // Miles
	kmEarthRadius = 6371 // Kilometers
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lng float64
}

func degreeToRadian(d float64) float64 {
	return d * math.Pi / 180
}

func Distance(p, q Coord) (mi, km float64) {
	lat1 := degreeToRadian(p.Lat)
	lng1 := degreeToRadian(p.Lng)
	lat2 := degreeToRadian(q.Lat)
	lng2 := degreeToRadian(q.Lng)

	diffLat := lat2 - lat1
	diffLng := lng2 - lng1

	a := math.Pow(math.Sin(diffLat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLng/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * miEarthRadius
	km = c * kmEarthRadius

	return mi, km
}
