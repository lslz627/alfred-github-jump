package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strings"
	"testing"
)

func TestUpdateRepositories(t *testing.T) {
	token := &oauth2.Token{}
	str := "{\"access_token\":\"3e51f128402058030266e916d0ada5f506262211\",\"token_type\":\"bearer\",\"expiry\":\"0001-01-01T00:00:00Z\"}\n"
	json.NewDecoder(strings.NewReader(str)).Decode(token)

	tc := OAuthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(tc)

	userRepos, _ := listUserRepositories(client)

	for _, repo := range userRepos {
		fmt.Printf("%+v\n", repo)
	}
}
