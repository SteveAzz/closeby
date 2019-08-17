package customers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/steveazz/closeby/pkg/geo"
)

func TestCustomer_InRangeCustomers(t *testing.T) {
	customers := []Customer{
		{
			ID:   12,
			Name: "Christina McArdle",
			Lat:  52.986375,
			Long: -6.043701,
		},
		{
			ID:   8,
			Name: "Eoin Ahearn",
			Lat:  54.0894797,
			Long: -6.18671,
		},
		{
			ID:   1,
			Name: "Alice Cahill",
			Lat:  51.92893,
			Long: -10.27699,
		},
	}

	expectedCustomers := []Customer{
		{
			ID:   12,
			Name: "Christina McArdle",
			Lat:  52.986375,
			Long: -6.043701,
		},
		{
			ID:   8,
			Name: "Eoin Ahearn",
			Lat:  54.0894797,
			Long: -6.18671,
		},
	}

	loc := geo.Location{
		Lat:  53.339428,
		Long: -6.257664,
	}

	inRangeCusomters := InRangeCustomers(customers, loc, 100.0)

	require.Equal(t, expectedCustomers, inRangeCusomters)
}
