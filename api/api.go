package api

import (
	"github.com/adamkasztenny/slack-repeat-bot/domain"
	"github.com/nlopes/slack"
	"strings"
)

const keyword = "say "

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

func (api *API) Listen() {
	go api.rtm.ManageConnection()

	for event := range api.rtm.GetIncomingEvents() {
		api.handleEvents(event)
	}
}

func (api *API) handleEvents(event slack.RTMEvent) {
	switch message := event.Data.(type) {
	case *slack.MessageEvent:
		api.handleIncomingMessage(message)
	}
}

func (api *API) handleIncomingMessage(incomingMessage *slack.MessageEvent) {
	if strings.HasPrefix(incomingMessage.Text, keyword) {
		api.sendMessage(incomingMessage)
	}
}

func (api *API) sendMessage(incomingMessage *slack.MessageEvent) {
	wordToRepeat := strings.Replace(incomingMessage.Text, keyword, "", 1)
	repeatedWord := domain.Repeat(wordToRepeat)
	api.rtm.SendMessage(api.rtm.NewOutgoingMessage(repeatedWord, incomingMessage.Channel))
}
