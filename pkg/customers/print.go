package customers

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
)

// PrintCustomers write to io.Writer with the help of a tab write to format the
// given customers. Only show the customer ID and the name, which is sorted by
// ID in ascending order.
func PrintCustomers(w io.Writer, customers []Customer) error {
	sort.Sort(ByID(customers))

	fmt.Fprint(w, "\tCLOSE BY CUSTOMERS\n\n")
	tab := tabwriter.NewWriter(w, 0, 2, 2, ' ', 0)
	for _, c := range customers {
		fmt.Fprintf(w, "\t%d \t%s\n", c.ID, c.Name)
	}
	fmt.Fprint(w, "\n")

	return tab.Flush()
}
