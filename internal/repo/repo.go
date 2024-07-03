package repo

import (
	"fmt"
	"strings"
)

const (
	GitHubBaseURL = "https://github.com/"
	GitHubGitExt  = ".git"
)

func ParseRepo(url string) (string, string, error) {
	path := strings.TrimPrefix(url, GitHubBaseURL)
	path = strings.TrimSuffix(path, GitHubGitExt)

	parts := strings.Split(path, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid repo URL format")
	}
	return parts[0], parts[1], nil
}
