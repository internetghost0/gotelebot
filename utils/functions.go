package utils

import (
	"os/exec"
)

func ExecCmd(cmd string, args []string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
