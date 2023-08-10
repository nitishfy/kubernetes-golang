package contexts

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetContexts() ([]string, error) {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	cmd := exec.Command("kubectl", "--kubeconfig", kubeconfigPath, "config", "get-contexts", "-o=name")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
		return nil, err
	}
	contexts := strings.Split(strings.TrimSpace(string(output)), "\n")
	return contexts, nil
}
