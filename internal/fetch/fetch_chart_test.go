package fetch_test

import (
	"helm-generator/internal/fetch"
	"os"
	"path/filepath"
	"testing"
)

func TestFetchChart(t *testing.T) {
	t.Run("Already exists", func(t *testing.T) {
		chartName := "test-chart"
		version := "1.0.0"
		dest := "/tmp/charts"
		helmPath := "/usr/local/bin/helm"
		chartDir := filepath.Join(dest, chartName)

		if err := os.Mkdir(chartDir, 0755); err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(chartDir)

		cfg := fetch.ChartConfig{
			Name:        chartName,
			Version:     version,
			Destination: dest,
			Path:        helmPath,
		}

		chart := fetch.NewChart(cfg)

		err := chart.Fetch()

		// Expect no error to be returned and message to be printed
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
	})

	t.Run("Fetch new chart", func(t *testing.T) {
		chartName := "mysql"
		version := "9.7.0"
		dest := "/tmp/charts"
		helmPath := "/usr/local/bin/helm"
		chartDir := filepath.Join(dest, chartName)

		// Remove chart directory if it exists
		os.RemoveAll(chartDir)

		// Create the parent directory first
		if err := os.MkdirAll(dest, 0755); err != nil {
			t.Fatal(err)
		}

		cfg := fetch.ChartConfig{
			Name:        chartName,
			Version:     version,
			Destination: dest,
			Path:        helmPath,
		}

		chart := fetch.NewChart(cfg)

		err := chart.Fetch()

		// Expect no error to be returned and chart directory to be created
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}

		if _, err := os.Stat(chartDir); os.IsNotExist(err) {
			t.Errorf("Expected chart directory %s to exist but it doesn't", chartDir)
		}

		// Clean up chart directory
		os.RemoveAll(chartDir)
	})
}
