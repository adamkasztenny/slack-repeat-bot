package api

import "github.com/nlopes/slack"

type RTMStub struct {
	incomingEvents chan slack.RTMEvent
}

func (RTMStub) ManageConnection() {
}

func (rtm RTMStub) NewOutgoingMessage(text string, channelID string, options ...slack.RTMsgOption) *slack.OutgoingMessage {
	return nil
}

func (rtm RTMStub) GetIncomingEvents() chan slack.RTMEvent {
	return rtm.incomingEvents
}

func (rtm RTMStub) SetIncomingEvents(events chan slack.RTMEvent) {
	rtm.incomingEvents = events
}
