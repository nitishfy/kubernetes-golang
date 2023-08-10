package kubectx

import (
	"os/exec"
	"strings"
)

func RunKubectx() ([]string, error) {
	cmd := exec.Command("kubectx")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	contexts := strings.Split(strings.TrimSpace(string(output)), "\n")
	return contexts, nil
}
