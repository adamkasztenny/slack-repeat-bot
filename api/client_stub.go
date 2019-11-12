package api

import "github.com/nlopes/slack"

type ClientStub struct {
	user      *slack.User
	userError error
}

func (clientStub *ClientStub) NewRTM() *slack.RTM {
	return nil
}

func (clientStub *ClientStub) GetUserInfo(userID string) (*slack.User, error) {
	return clientStub.user, clientStub.userError
}

func (clientStub *ClientStub) SetUserInfo(user *slack.User, userError error) {
	clientStub.user = user
	clientStub.userError = userError
}
