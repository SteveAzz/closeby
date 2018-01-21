package customers

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomer_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name        string
		data        string
		expected    Customer
		expectedErr string
	}{
		{
			name: "Valid Customer",
			data: `{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`,
			expected: Customer{
				ID:   12,
				Name: "Christina McArdle",
				Lat:  52.986375,
				Long: -6.043701,
			},
		},
		{
			name:        "Invalid customer id",
			data:        `{"latitude": "52.986375", "user_id": "INVALID_ID", "name": "Christina McArdle", "longitude": "-6.043701"}`,
			expectedErr: "failed to type assert ID to int type: string",
		},
		{
			name:        "Invalid latitude",
			data:        `{"latitude": "INVALID", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`,
			expectedErr: "failed to transform latitude to float64 type: string: strconv.ParseFloat: parsing \"INVALID\": invalid syntax",
		},
		{
			name:        "Invalid longitude",
			data:        `{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "INVALID"}`,
			expectedErr: "failed to transform longitude to float64 type: string: strconv.ParseFloat: parsing \"INVALID\": invalid syntax",
		},
		{
			name:        "Invalid name",
			data:        `{"latitude": "52.986375", "user_id": 12, "name": 222, "longitude": "-6.043701"}`,
			expectedErr: "failed to type assert Name to int type: float64",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cst := Customer{}
			err := json.Unmarshal([]byte(c.data), &cst)
			if c.expectedErr != "" {
				require.EqualError(t, err, c.expectedErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, c.expected, cst)
		})
	}
}

func TestCustomer_ReadFromFile(t *testing.T) {
	cases := []struct {
		name        string
		loc         string
		expected    []Customer
		expectedErr bool
	}{
		{
			name: "File expectedErr",
			loc:  "./fixtures/customers.txt",
			expected: []Customer{
				{
					ID:   12,
					Name: "Christina McArdle",
					Lat:  52.986375,
					Long: -6.043701,
				},
				{
					ID:   1,
					Name: "Alice Cahill",
					Lat:  51.92893,
					Long: -10.27699,
				},
			},
		},
		{
			name:        "File does not exist",
			loc:         "randomlocation",
			expectedErr: true,
		},
		{
			name:        "File is not valid json",
			loc:         "./fixtures/invalid.txt",
			expectedErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			customers, err := ReadFromFile(c.loc)
			if c.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, c.expected, customers)
		})
	}
}
