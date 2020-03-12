package haversine

import (
	"fmt"
	"testing"
)

var tests = []struct {
	p     Coord
	q     Coord
	miles float64
	kms   float64
}{
	{
		Coord{Lat: 38.898556, Lng: -77.037852}, // 1600 Pennsylvania Ave NW, Washington, DC
		Coord{Lat: 38.897147, Lng: -77.043934}, // 1922 F St NW, Washington, DC
		0.341,
		0.549,
	},
	{
		Coord{Lat: 13.0103662, Lng: 80.2574283}, // Theosophical Soceity Adyar
		Coord{Lat: 8.0781423, Lng: 77.5531069},  // Vivekananda Memorial Kanniyakumari
		387.036,
		622.994,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		mi, km := Distance(input.p, input.q)
		if fmt.Sprintf("%.3f", input.miles) != fmt.Sprintf("%.3f", mi) || fmt.Sprintf("%.3f", input.kms) != fmt.Sprintf("%.3f", km) {
			t.Errorf("fail: want %v %v -> %v %v got %v %v",
				input.p,
				input.q,
				fmt.Sprintf("%.3f", input.miles),
				fmt.Sprintf("%.3f", input.kms),
				fmt.Sprintf("%.3f", mi),
				fmt.Sprintf("%.3f", km),
			)
		}
	}
}
