package main

import (
	"eth_usd/csv_process"
	"eth_usd/plotter"
	"fmt"
	"gonum.org/v1/plot/vg"
	"os"
)

func main() {
	pricePairs, err := csvprocess.LoadDataFrom("prices.csv")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading prices from csv file: %v\n", err)
		os.Exit(1)
	}

	pricePlots, err := plotter.GeneratePlotFor(pricePairs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building chart: %v\n", err)
		os.Exit(1)
	}
	if err := pricePlots.Save(15*vg.Inch, 4*vg.Inch, "ethereum_pirces.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving plot: %v\n", err)
		os.Exit(1)
	}
}
