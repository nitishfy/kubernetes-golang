package namespace

import (
	"log"
	"os/exec"
	"strings"
)

func GetNamespace() ([]string, error) {
	cmd := exec.Command("kubectl", "get", "namespace")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
		return nil, err
	}
	namespace := strings.Split(strings.TrimSpace(string(output)), "\n")
	return namespace, err
}
