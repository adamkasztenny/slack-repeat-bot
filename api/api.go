package api

import (
	"github.com/adamkasztenny/slack-repeat-bot/domain"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
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
	log.Info("Listening...")
	go api.rtm.ManageConnection()

	for event := range api.rtm.GetIncomingEvents() {
		log.Debugf("Got event: %v", event)
		api.handleEvents(event)
	}
}

func (api *API) handleEvents(event slack.RTMEvent) {
	switch message := event.Data.(type) {
	case *slack.ConnectedEvent:
		log.Info("Connected successfully.")
	case *slack.MessageEvent:
		api.handleIncomingMessage(message)
	case *slack.RTMError:
		log.Errorf("Error: %s\n", message.Error())
	}
}

func (api *API) handleIncomingMessage(incomingMessage *slack.MessageEvent) {
	if strings.HasPrefix(incomingMessage.Text, keyword) {
		log.Infof("Sending repeated message to channel: %v", incomingMessage.Channel)
		api.sendMessage(incomingMessage)
	}
}

func (api *API) sendMessage(incomingMessage *slack.MessageEvent) {
	wordToRepeat := strings.Replace(incomingMessage.Text, keyword, "", 1)
	repeatedWord := domain.Repeat(wordToRepeat)
	api.rtm.SendMessage(api.rtm.NewOutgoingMessage(repeatedWord, incomingMessage.Channel))
}
