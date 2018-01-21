package customers

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// Customer represents information about a specific customer.
type Customer struct {
	ID   int     `json:"user_id"`
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

// UnmarshalJSON implements the Unmarshaler interface to do custom unmarshaling
// on the customers data.
func (c *Customer) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	for k, v := range data {
		switch k {
		case "user_id":
			id, ok := v.(float64)
			if !ok {
				return errors.Errorf("failed to type assert ID to int type: %T", v)
			}
			c.ID = int(id)
		case "name":
			var ok bool
			if c.Name, ok = v.(string); !ok {
				return errors.Errorf("failed to type assert Name to int type: %T", v)
			}
		case "latitude":
			f, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return errors.Wrapf(err, "failed to transform latitude to float64 type: %T", v)
			}
			c.Lat = f
		case "longitude":
			f, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return errors.Wrapf(err, "failed to transform longitude to float64 type: %T", v)
			}
			c.Long = f
		}
	}

	return nil
}

// ReadFromFile reads a file with one customer per line, JSON-encoded and
// returns a slice of customers. Only loads small files with this since it loads
// its content in memory. Returns an error if it fails to load the file or read
// the data in that file.
func ReadFromFile(loc string) ([]Customer, error) {
	var err error
	var customers []Customer

	f, err := os.Open(loc)
	defer func() {
		if err = f.Close(); err != nil {
			err = errors.Wrap(err, "close on customer file failed")
		}
	}()
	if err != nil {
		return nil, errors.Wrap(err, "read on customer file failed")
	}

	fileScanner := bufio.NewScanner(f)

	// For each line marshal into a customer.
	for fileScanner.Scan() {
		c := Customer{}
		err = json.Unmarshal(fileScanner.Bytes(), &c)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal customer data")
		}

		customers = append(customers, c)
	}

	return customers, err
}
