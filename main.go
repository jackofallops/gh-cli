package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jackofallops/gh-cli/cmd"
	"github.com/jackofallops/gh-cli/utils"

)

func main() {
	owner := "terraform-providers"
	repo := "terraform-provider-azurerm"

	prCmd := flag.NewFlagSet("pr", flag.ExitOnError)
	activityCmd := flag.NewFlagSet("activity", flag.ExitOnError)

	prCmdRepo := prCmd.String("repo", repo, "Repo to filter for activity. (Required)")
	prCmdOwner := prCmd.String("owner", owner, "Org or Owner of the repo to filter on. (Required)")
	prCmdCount := prCmd.Int("count", 200, "Max number of PR's to return. Defaults to 200")

	activityCmdRepo := activityCmd.String("repo", repo, "Repo to filter for activity. (Required)")
	activityCmdOwner := activityCmd.String("owner", owner, "Org or Owner of the repo to filter on. (Required)")
	activityCmdLogin := activityCmd.String("login", "", "User to query for. (Required)")

	if len(os.Args) < 2 {
		fmt.Println("Error: subcommand is required (`pr` or `activity`)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "pr":
		prCmd.Parse(os.Args[2:])
	case "activity":
		activityCmd.Parse(os.Args[2:])
		if *activityCmdLogin == "" {
			activityCmd.PrintDefaults()
			os.Exit(2)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if prCmd.Parsed() {
		prs := cmd.List(*prCmdOwner, *prCmdRepo, *prCmdCount)
		fmt.Println("number§title§createdAt§user§url")
		for _, pr := range prs {
			fmt.Printf("%d§%s§%s§%s§%s\n",
				*pr.Number,
				utils.DerefStringSafely(pr.Title),
				pr.CreatedAt.String(),
				*pr.User.Login,
				*pr.HTMLURL,
			)
		}
	}

	if activityCmd.Parsed() {
		fmt.Printf("Activity: %+v", cmd.GetUserActivityForRepo(*activityCmdLogin, *activityCmdOwner, *activityCmdRepo))
	}

}


