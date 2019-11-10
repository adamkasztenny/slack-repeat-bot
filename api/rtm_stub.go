package api

import "github.com/nlopes/slack"

type RTMStub struct {
	incomingEvents    chan slack.RTMEvent
	mostRecentMessage *slack.OutgoingMessage
}

func (rtm *RTMStub) ManageConnection() {
}

func (rtm *RTMStub) NewOutgoingMessage(text string, channelID string, options ...slack.RTMsgOption) *slack.OutgoingMessage {
	return &slack.OutgoingMessage{Text: text, Channel: channelID}
}

func (rtm *RTMStub) GetIncomingEvents() chan slack.RTMEvent {
	return rtm.incomingEvents
}

func (rtm *RTMStub) SetIncomingEvents(events chan slack.RTMEvent) {
	rtm.incomingEvents = events
}

func (rtm *RTMStub) SendMessage(msg *slack.OutgoingMessage) {
	rtm.mostRecentMessage = msg
}

func (rtm *RTMStub) GetMostRecentMessage() *slack.OutgoingMessage {
	return rtm.mostRecentMessage
}
