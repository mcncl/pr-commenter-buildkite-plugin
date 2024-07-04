package repo

import (
	"fmt"
	"strings"
)

const (
	GitHubBaseHTTPS = "https://github.com/"
	GitHubBaseSSH   = "git@github.com:"
	GitHubGitExt    = ".git"
)

func ParseRepo(url string) (string, string, error) {
	var path string
	if strings.HasPrefix(url, GitHubBaseHTTPS) {
		path = strings.TrimPrefix(url, GitHubBaseHTTPS)
		path = strings.TrimSuffix(path, GitHubGitExt)
	} else if strings.HasPrefix(url, GitHubBaseSSH) {
		path = strings.TrimPrefix(url, GitHubBaseSSH)
		path = strings.TrimSuffix(path, GitHubGitExt)
	} else {
		return "", "", fmt.Errorf("invalid repo URL format")
	}

	parts := strings.Split(path, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid repo URL format")
	}
	return parts[0], parts[1], nil
}
