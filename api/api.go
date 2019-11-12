package api

import (
	"fmt"
	"github.com/adamkasztenny/slack-repeat-bot/domain"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

const keyword = "say "

var keywordRegex = regexp.MustCompile(keyword + "(.*)")

type API struct {
	client ClientInterface
	rtm    RTMInterface
}

func (api *API) Initialize(token string) {
	api.client = Client{
		client: slack.New(token),
	}
	api.rtm = RTM{
		rtm: api.client.NewRTM(),
	}
}

func (api *API) Listen() {
	log.Info("Connecting...")
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
	if keywordRegex.MatchString(incomingMessage.Text) {
		log.Infof("Sending repeated message to channel: %v", incomingMessage.Channel)
		api.sendMessage(incomingMessage)
	}
}

func (api *API) sendMessage(incomingMessage *slack.MessageEvent) {
	username := api.getUsername(incomingMessage)
	message := api.getMessage(incomingMessage, username)
	api.rtm.SendMessage(api.rtm.NewOutgoingMessage(message, incomingMessage.Channel))
}

func (api *API) getUsername(incomingMessage *slack.MessageEvent) string {
	user, err := api.client.GetUserInfo(incomingMessage.User)

	if err != nil {
		log.Errorf("Cannot get user with id %v: %v", incomingMessage.User, err)
		return ""
	}
	return user.Name
}

func (api *API) getMessage(incomingMessage *slack.MessageEvent, username string) string {
	wordToRepeat := strings.Replace(incomingMessage.Text, keyword, "", 1)
	repeatedWord := domain.Repeat(wordToRepeat)
	return fmt.Sprintf("%s says: %s", username, repeatedWord)
}
