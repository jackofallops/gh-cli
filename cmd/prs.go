package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"os"
)

func List(owner, repo string, count int) []*github.PullRequest {
	client := github.NewClient(nil)
	listOpts := github.ListOptions{PerPage: count}
	prs, resp, err := client.PullRequests.List(context.Background(), owner, repo, &github.PullRequestListOptions{ListOptions: listOpts})
	if err != nil {
		fmt.Printf("response: %d, error: %+v", resp.StatusCode, err)
		os.Exit(3)
	}

	return prs

}

