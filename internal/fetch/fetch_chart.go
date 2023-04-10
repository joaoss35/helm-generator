package fetch

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func FetchChart(chartName *string, version *string, dest *string, helmPath *string) error {
	// Check if chart already exists in destination directory
	chartDir := filepath.Join(*dest, *chartName)
	if _, err := os.Stat(chartDir); err == nil {
		fmt.Printf("The '%s' chart already exists in directory '%s', skipping fetch...\n", *chartName, *dest)
		return nil
	}

	cmdArgs := []string{"pull", fmt.Sprintf("bitnami/%s", *chartName), "--destination", *dest, "--untar"}

	// Only add the version flag if the version is not empty
	if *version != "" {
		cmdArgs = append(cmdArgs, "--version", *version)
	}

	cmd := exec.Command(*helmPath, cmdArgs...)
	errExec := cmd.Run()
	if errExec != nil {
		log.Fatal(errExec)
	}

	fmt.Printf("Successfully fetched chart '%s' version '%s' to directory '%s'\n", *chartName, *version, *dest)
	return nil
}
