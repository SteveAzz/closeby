package customers

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update .golden files")

func TestCustomer_PrintCustomer(t *testing.T) {
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

	// Update contents of golden file if update flag is set to true, generally
	// this is done by `go test github.com/SteveAzz/closeby/pkg/customers -update`
	if *update {
		f, err := os.Create("./fixtures/print.golden")
		defer func(t *testing.T) {
			err = f.Close()
			require.NoError(t, err)
		}(t)
		require.NoError(t, err)

		err = PrintCustomers(f, customers)
		require.NoError(t, err)
	}

	// Get golden file contents.
	expected, err := ioutil.ReadFile("./fixtures/print.golden")
	require.NoError(t, err)

	// Write to string.
	got := new(bytes.Buffer)
	err = PrintCustomers(got, customers)
	require.NoError(t, err)

	require.Equal(t, string(expected), got.String())
}
