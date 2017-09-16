package main

import (
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"io/ioutil"
	"os"
	"strings"
)

var org = flag.String("org", "", "the organization to match against")

func main() {
	flag.Parse()

	if *org == "" {
		fmt.Println("Error:\n  -org is required")
		os.Exit(1)
	}

	fmt.Println("Connecting to Github...")

	tokenbytes, err := ioutil.ReadFile(".token")
	token := string(tokenbytes)

	if err != nil {
		fmt.Printf("Failed to read token file: %v", err)
		panic(err)
	}

	client := NewGithubClient(token)

	user := client.GetUser()
	allOrgs := client.GetOrgs()

	fmt.Printf("Got user: %q\n", user.GetLogin())
	fmt.Println()
	fmt.Printf("%q is member of %q? %v\n", user.GetLogin(), *org, ValidateOrgMembership(allOrgs, *org))
}

func ValidateOrgMembership(allOrgs []*github.Organization, org string) bool {
	for _, o := range allOrgs {
		if strings.ToLower(*o.Login) == strings.ToLower(org) {
			return true
		}
	}
	return false
}
