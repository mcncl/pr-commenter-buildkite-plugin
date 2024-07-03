package secret

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ExecCommand = exec.Command

func GetSecret(name string) (string, error) {
	if name == "GITHUB_TOKEN" {
		_, found := os.LookupEnv(name)
		if !found {
			return "", fmt.Errorf("could not use %s, secret does not exist on cluster", name)
		}
	}
	cmd := ExecCommand("buildkite-agent", "secret", "get", name)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
