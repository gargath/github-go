package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"
	"os"
)

type githubClient struct {
	client *github.Client
}

func NewGithubClient(token string) *githubClient {
	tc := cleanhttp.DefaultClient()
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, tc)
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc = oauth2.NewClient(ctx, ts)
	myClient := github.NewClient(tc)

	c := &githubClient{client: myClient}
	return c
}

func (c *githubClient) GetUser() *github.User {
	user, _, err := c.client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("Failed to talk to Github: %q\n", err)
		os.Exit(1)
	}
	return user
}

func (c *githubClient) GetOrgs() []*github.Organization {
	orgOpt := &github.ListOptions{
		PerPage: 100,
	}

	var allOrgs []*github.Organization
	for {
		orgs, resp, err := c.client.Organizations.List(context.Background(), "", orgOpt)
		if err != nil {
			fmt.Printf("Failed to get Organizations: %q\n", err)
			os.Exit(1)
		}
		allOrgs = append(allOrgs, orgs...)
		if resp.NextPage == 0 {
			break
		}
		orgOpt.Page = resp.NextPage
	}
	return allOrgs
}
