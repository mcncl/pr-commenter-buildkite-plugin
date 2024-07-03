package secret_test

import (
	"os/exec"
	"testing"

	"prcommenter/internal/secret"
)

func TestGetSecret(t *testing.T) {
	oldExecCmd := secret.ExecCommand
	defer func() { secret.ExecCommand = oldExecCmd }()

	secret.ExecCommand = func(command string, args ...string) *exec.Cmd {
		return exec.Command("echo", "foobar")
	}

	got, err := secret.GetSecret("MOCK_SECRET")
	if err != nil {
		t.Fatalf("error getting secret value: %s", err)
	}

	want := "foobar"
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}
