package api

import "github.com/nlopes/slack"

type RTMInterface interface {
	GetIncomingEvents() chan slack.RTMEvent
	ManageConnection()
	NewOutgoingMessage(text string, channelID string, options ...slack.RTMsgOption) *slack.OutgoingMessage
}

type RTM struct {
	rtm *slack.RTM
}

func (rtm RTM) GetIncomingEvents() chan slack.RTMEvent {
	return rtm.rtm.IncomingEvents
}

func (rtm RTM) ManageConnection() {
	rtm.rtm.ManageConnection()
}

func (rtm RTM) NewOutgoingMessage(text string, channelID string, options ...slack.RTMsgOption) *slack.OutgoingMessage {
	return rtm.rtm.NewOutgoingMessage(text, channelID, options...)
}
