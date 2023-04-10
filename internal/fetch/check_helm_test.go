package fetch_test

import (
	"helm-generator/internal/fetch"
	"os"
	"testing"
)

func TestCheckHelm(t *testing.T) {
	// Test when Helm is not installed
	err := os.Setenv("PATH", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Unsetenv("PATH")

	// Expect CheckHelm to return an error
	err = fetch.CheckHelm()
	if err == nil {
		t.Errorf("CheckHelm did not return an error")
	}
}

func TesCheckHelmWithHelmInstalled(t *testing.T) {
	// Test when Helm is installed
	err := os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Unsetenv("PATH")

	// Expect CheckHelm to not return an error
	err = fetch.CheckHelm()
	if err != nil {
		t.Errorf("CheckHelm returned an error: %v", err)
	}
}
