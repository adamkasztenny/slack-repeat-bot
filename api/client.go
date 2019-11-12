package api

import "github.com/nlopes/slack"

type ClientInterface interface {
	NewRTM() *slack.RTM
	GetUserInfo(userID string) (*slack.User, error)
}

type Client struct {
	client *slack.Client
}

func (client Client) NewRTM() *slack.RTM {
	return client.client.NewRTM()
}

func (client Client) GetUserInfo(userID string) (*slack.User, error) {
	return client.client.GetUserInfo(userID)
}
