package customers

import (
	"github.com/steveazz/closeby/pkg/geo"
)

// InRangeCustomers takes a list of customers and checks if it is within the
// given range for the specified location.
func InRangeCustomers(customers []Customer, loc geo.Location, radius float64) (inRangeCustomers []Customer) {
	for _, c := range customers {
		dist := loc.GreatCircleDistance(&geo.Location{Lat: c.Lat, Long: c.Long})

		if dist <= radius {
			inRangeCustomers = append(inRangeCustomers, c)
		}
	}

	return
}
