package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
)

func GetUserActivityForRepo(login, owner, repo string) (userEvents []string) {
	client := github.NewClient(nil)
	listOpts := &github.ListOptions{PerPage: 10000}
	events, resp, err := client.Activity.ListEventsPerformedByUser(context.Background(), login, true, listOpts)
	if err != nil {
		fmt.Printf("Error getting response for user %s (%d), err: %+v", login, resp.StatusCode, err)
	}

	for _, event := range events {
		if *event.Repo.Name != (owner + "/" + repo) {
			fmt.Printf("repo %s does not match %s, skipping\n", *event.Repo.Name, repo)
			continue
		}
		eventPayload, err := event.ParsePayload()
		if err != nil {
			continue
		}
		userEventDecoded := fmt.Sprintf("type: %s, Content: %+v", *event.Type, eventPayload)
		userEvents = append(userEvents, userEventDecoded)
	}

	return
}
