package comment_test

import (
	"context"
	"testing"

	"prcommenter/internal/issue/comment"

	"github.com/google/go-github/github"
)

type mockGitHubClient struct {
	createComment func(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
}

func (m *mockGitHubClient) CreateComment(ctx context.Context, owner string, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	return m.createComment(ctx, owner, repo, number, comment)
}

func TestPost(t *testing.T) {
	mockClient := &mockGitHubClient{
		createComment: func(ctx context.Context, owner, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
			if owner != "testdev" || repo != "hello" || number != 420 || *comment.Body != "Test comment" {
				t.Errorf("Unexpected arguments to CreateComment")
			}
			return nil, nil, nil
		},
	}

	err := comment.Post(context.Background(), mockClient, "testdev", "hello", "420", "Test comment")
	if err != nil {
		t.Fatalf("error posting comment: %s", err)
	}
}

func TestPostCommentEmptyBody(t *testing.T) {
	mockClient := &mockGitHubClient{}

	err := comment.Post(context.Background(), mockClient, "testdev", "hello", "69", "")
	if err == nil {
		t.Fatalf("error expected due to empty body")
	}
}
