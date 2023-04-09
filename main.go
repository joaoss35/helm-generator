package main

import (
	"flag"
	"fmt"
	"helm-generator/fetch"
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
	helmErr := fetch.CheckHelm()

	// Check if Bitnami repository is added
	bitnamiErr := fetch.AddBitnamiRepo()

	if helmErr != nil {
		fmt.Println("Helm could not be found in PATH")
		os.Exit(1)
	}

	if bitnamiErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", bitnamiErr)
		os.Exit(1)
	}

	// Fetch the Helm path
	helmPath, pathErr := exec.LookPath("helm")
	if pathErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", pathErr)
		os.Exit(1)
	}

	// Fetch the chart
	err := fetch.FetchChart(chartName, version, dest, &helmPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
