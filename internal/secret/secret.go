package secret

import (
	"fmt"
	"os/exec"
	"strings"
)

var ExecCommand = exec.Command

func GetSecret(name string) (string, error) {
	cmd := ExecCommand("buildkite-agent", "secret", "get", name)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
