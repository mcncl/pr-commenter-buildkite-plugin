# Issue Commenter

The internal workings of the commenter plugin.

File structure:
```sh
.
├── common
│   ├── common.go
│   └── common_test.go
├── github
│   ├── client.go
│   └── client_test.go
├── issue
│   └── comment
│       ├── comment.go
│       └── comment_test.go
├── repo
│   ├── repo.go
│   └── repo_test.go
└── secret
    ├── secret.go
    └── secret_test.go
```

## common
Houses the prefix for the Buildkite plugin. Offers no other function right now.

## github
Responsible for creating the client that interacts with the GitHub API. A call to `github.New([token])` will create a new authenticated client with GitHub in the comment API.

## issue
The core logic for `POST`ing a comment to a GitHub issue. An Issue is an *easier* thing to post to as the comment becomes a part of the conversation rather than a review of a file change. The intent of this plugin is just that; outputting a `message` to the PR conversation.

NB: By the nature of GitHub Issues and PRs being able to use the same endpoint, this plugin *can* post to an Issue too.

## repo
The parsing login for a GitHub repository. This allows the plugin to support posting to a PR whether an SSH *or* HTTPS repository address has been provided.

## secret
The `GETTER` for the secret that should be used to authenticate with the GitHub API. The default value, if no `secret-name` is provided is `GITHUB_TOKEN`. The `buildkite-agent` API will be used to `GET` a secret that is available on a cluster.

⚠️ If clusters are not being used then this command will fail as secrets are not available on *unclustered* agents.
