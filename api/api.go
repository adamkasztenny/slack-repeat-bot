package api

import "github.com/nlopes/slack"

type API struct {
	client *slack.Client
	rtm    RTMInterface
}

func (api *API) Initialize(token string) {
	api.client = slack.New(token)
	api.rtm = RTM{
		rtm: api.client.NewRTM(),
	}
}
