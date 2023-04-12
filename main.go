package main

import (
	"flag"
	"fmt"
	"helm-generator/internal/fetch"
	"os"
	"os/exec"
)

func main() {
	// Parse flags
	chartName := flag.String("chart", "", "name of the Helm chart to fetch")
	version := flag.String("version", "", "version of the Helm chart to fetch")
	dest := flag.String("dest", ".", "destination directory to save the Helm chart")
	help := flag.Bool("help", false, "display help message")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	// Check if Helm is installed
	err := fetch.CheckHelm()

	if err != nil {
		fmt.Println("Helm could not be found in PATH")
		os.Exit(1)
	}

	// Check if Bitnami repository is added
	err = fetch.AddBitnamiRepo()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Fetch the Helm path
	helmPath, err := exec.LookPath("helm")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cfg := fetch.ChartConfig{
		Name:        *chartName,
		Version:     *version,
		Destination: *dest,
		Path:        helmPath,
	}

	chart := fetch.NewChart(cfg)

	// Fetch the chart
	err = chart.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
