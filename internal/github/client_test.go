package github_test

import (
	"prcommenter/internal/github"
	"testing"
)

func TestNew(t *testing.T) {
	client, err := github.New("mock")
	if err != nil {
		t.Fatalf("error creating client: %s", err)
	}
	if client == nil {
		t.Fatal("expected client to be returned")
	}
}
