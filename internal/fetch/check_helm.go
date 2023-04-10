package fetch

import (
	"os/exec"
)

func CheckHelm() error {
	_, err := exec.LookPath("helm")
	if err != nil {
		return err
	}
	return nil
}
