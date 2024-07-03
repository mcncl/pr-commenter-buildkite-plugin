package repo_test

import (
	"prcommenter/internal/repo"
	"testing"
)

func TestParseRepoInfo(t *testing.T) {
	tests := []struct {
		name      string
		repoURL   string
		wantOwner string
		wantRepo  string
		wantErr   bool
	}{
		{
			name:      "Valid GitHub URL",
			repoURL:   "https://github.com/gooddev/hello.git",
			wantOwner: "gooddev",
			wantRepo:  "hello",
			wantErr:   false,
		},
		{
			name:      "Invalid URL format",
			repoURL:   "https://github.com/invalid-format",
			wantOwner: "",
			wantRepo:  "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOwner, gotRepo, err := repo.ParseRepo(tt.repoURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("wanted error %v, got %v", tt.wantErr, err)
				return
			}
			if gotOwner != tt.wantOwner {
				t.Errorf("wanted owner %s, got %s", tt.wantOwner, gotOwner)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("wanted repo %s, got %s", tt.wantRepo, gotRepo)
			}
		})
	}
}
