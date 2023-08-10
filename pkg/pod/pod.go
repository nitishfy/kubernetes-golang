package pod

import (
	"log"
	"os/exec"
	"strings"
)

func GetPods() ([]string, error) {
	cmd := exec.Command("kubectl", "get", "pods")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
		return nil, err
	}
	pod := strings.Split(strings.TrimSpace(string(output)), "\n")
	return pod, nil
}
