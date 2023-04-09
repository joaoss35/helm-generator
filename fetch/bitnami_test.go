package fetch_test

import (
	"bytes"
	"os/exec"
	"testing"

	"helm-generator/fetch"
)

func TestBitnamiRepoExists(t *testing.T) {
	// Check if bitnami repo is already added
	repoList, err := exec.Command("helm", "repo", "list").Output()
	if err != nil {
		t.Fatalf("Error running `helm repo list`: %v", err)
	}
	if bytes.Contains(repoList, []byte("bitnami")) {
		t.Skipf("bitnami repo already exists")
	}
}

func TestAddBitnamiRepo(t *testing.T) {
	// Add bitnami repo
	err := fetch.AddBitnamiRepo()
	if err != nil {
		t.Fatalf("Error adding bitnami repo: %v", err)
	}
}

func TestBitnamiRepoAdded(t *testing.T) {
	// Check if bitnami repo is added
	repoList, err := exec.Command("helm", "repo", "list").Output()
	if err != nil {
		t.Fatalf("Error running `helm repo list`: %v", err)
	}
	if !bytes.Contains(repoList, []byte("bitnami")) {
		t.Errorf("bitnami repo not added")
	}
}
