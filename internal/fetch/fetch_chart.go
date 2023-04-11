package fetch

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type ChartConfig struct {
	Name        string
	Version     string
	Destination string
	Path        string
}

type Chart struct {
	Config ChartConfig
}

func NewChart(cfg ChartConfig) *Chart {
	chart := &Chart{
		Config: cfg,
	}

	return chart
}

// func FetchChart(cfg ChartConfig) error {
func (chart *Chart) Fetch() error {
	// Check if chart already exists in destination directory
	chartDir := filepath.Join(chart.Config.Destination, chart.Config.Name)
	if _, err := os.Stat(chartDir); err == nil {
		fmt.Printf("The '%s' chart already exists in directory '%s', skipping fetch...\n", chart.Config.Name, chart.Config.Destination)
		return nil
	}

	cmdArgs := []string{"pull", fmt.Sprintf("bitnami/%s", chart.Config.Name), "--destination", chart.Config.Destination, "--untar"}

	// Only add the version flag if the version is not empty
	if chart.Config.Version != "" {
		cmdArgs = append(cmdArgs, "--version", chart.Config.Version)
	}

	cmd := exec.Command(chart.Config.Path, cmdArgs...)
	errExec := cmd.Run()
	if errExec != nil {
		log.Fatal(errExec)
	}

	fmt.Printf("Successfully fetched chart '%s' version '%s' to directory '%s'\n", chart.Config.Name, chart.Config.Version, chart.Config.Destination)
	return nil
}
