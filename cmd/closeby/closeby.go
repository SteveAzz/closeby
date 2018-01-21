package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/SteveAzz/closeby/pkg/customers"
)

func main() {
	fs := flag.NewFlagSet("closebycli", flag.ExitOnError)
	var (
		lat  = fs.String("lat", "53.339428", "Latitude of the location.")
		long = fs.String("long", "-6.257664", "Longitude of the location.")
		loc  = fs.String("c", "", "Location of the list of c.")
	)
	fs.Usage = usageOf(fs, os.Args[0]+" -c $FILELOCATION")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, "failed to parse flags")
		fs.Usage()
		os.Exit(1)
	}

	// If no customer list is shown fail to start the cli.
	if *loc == "" {
		fs.Usage()
		os.Exit(1)
	}

	lstOfCst, err := customers.ReadFromFile(*loc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	log.Printf("customers: %#+v", lstOfCst)
	log.Printf("lat: %#+v", *lat)
	log.Printf("long: %#+v", *long)
}

func usageOf(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprint(os.Stdout, "USAGE\n")
		fmt.Fprintf(os.Stdout, "  %s\n", short)
		fmt.Fprint(os.Stdout, "\n")
		fmt.Fprint(os.Stdout, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		if err := w.Flush(); err != nil {
			os.Exit(1)
		}
		fmt.Fprint(os.Stdout, "\n")
	}
}
