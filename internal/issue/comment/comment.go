package comment

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/go-github/github"
)

type GitHubClient interface {
	CreateComment(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
}

func Post(ctx context.Context, client GitHubClient, owner string, repo string, number string, body string) error {
	numberConverted, err := strconv.Atoi(number)
	if err != nil {
		return err
	}

	if body == "" {
		return errors.New("no body provided for comment")
	}

	comment := &github.IssueComment{
		Body: &body,
	}

	_, _, err = client.CreateComment(ctx, owner, repo, numberConverted, comment)
	return err
}
