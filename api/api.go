package api

import "github.com/nlopes/slack"

type API struct {
	client *slack.Client
	rtm    *slack.RTM
}

func (api *API) Initialize(token string) {
	api.client = slack.New(token)
	api.rtm = api.client.NewRTM()
}
