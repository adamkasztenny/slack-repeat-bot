package api

import "github.com/nlopes/slack"

type API struct {
	client *slack.Client
}

func (api *API) Initialize(token string) {
	api.client = slack.New(token)
}
