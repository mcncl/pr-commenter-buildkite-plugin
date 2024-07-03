package github

import (
	"context"

	"prcommenter/internal/issue/comment"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func New(token string) (comment.GitHubClient, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	return client.Issues, nil
}
