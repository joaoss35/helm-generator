package fetch

import (
	"bytes"
	"log"
	"os/exec"
)

func AddBitnamiRepo() error {
	repoList, err := exec.Command("helm", "repo", "list").Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if !bytes.Contains(repoList, []byte("bitnami")) {
		cmd := exec.Command("helm", "repo", "add", "bitnami", "https://charts.bitnami.com/bitnami")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
