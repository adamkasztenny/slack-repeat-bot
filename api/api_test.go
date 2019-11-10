package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type APITestSuite struct {
	suite.Suite
}

func (suite *APITestSuite) SetupTest() {
	gofakeit.Seed(0)
}

func (suite *APITestSuite) TestInitialization() {
	api := new(API)
	token := gofakeit.UUID()

	api.Initialize(token)

	assert.NotZero(suite.T(), api)
	assert.NotNil(suite.T(), api.client)
	assert.NotNil(suite.T(), api.rtm)
}

func (suite *APITestSuite) TestConnectsSuccessfully() {
	api := new(API)
	rtm := suite.createRTM(api)

	rtm.incomingEvents <- slack.RTMEvent{
		Data: &slack.ConnectedEvent{},
	}
	close(rtm.incomingEvents)

	assert.NotPanics(suite.T(), func() {
		api.Listen()
	})

	message := rtm.GetMostRecentMessage()
	assert.Nil(suite.T(), message)
}

func (suite *APITestSuite) TestSendsAMessageWithRepeatedTextIfTheKeywordIsPresent() {
	api := new(API)
	rtm := suite.createRTM(api)

	word := "word"
	channel := gofakeit.UUID()
	suite.sendMessage(rtm, keyword+word, channel)

	api.Listen()

	message := rtm.GetMostRecentMessage()
	assert.NotNil(suite.T(), message)
	assert.Equal(suite.T(), channel, message.Channel)
	assert.Contains(suite.T(), message.Text, word)
	assert.Contains(suite.T(), message.Text, "wowo")
	assert.NotContains(suite.T(), message.Text, keyword)
}

func (suite *APITestSuite) TestDoesNotSendAMessageIfTheKeywordIsNotPresent() {
	api := new(API)
	rtm := suite.createRTM(api)

	word := "word"
	channel := gofakeit.UUID()
	suite.sendMessage(rtm, word, channel)

	api.Listen()

	message := rtm.GetMostRecentMessage()
	assert.Nil(suite.T(), message)
}

func (suite *APITestSuite) TestHandlesError() {
	api := new(API)
	rtm := suite.createRTM(api)

	rtm.incomingEvents <- slack.RTMEvent{
		Data: &slack.RTMError{},
	}
	close(rtm.incomingEvents)

	assert.NotPanics(suite.T(), func() {
		api.Listen()
	})
	message := rtm.GetMostRecentMessage()
	assert.Nil(suite.T(), message)
}

func (suite *APITestSuite) createRTM(api *API) *RTMStub {
	rtmStub := new(RTMStub)
	incomingEvents := make(chan slack.RTMEvent, 1)
	rtmStub.SetIncomingEvents(incomingEvents)
	api.rtm = rtmStub
	return rtmStub

}

func (suite *APITestSuite) sendMessage(rtm *RTMStub, text, channel string) {
	rtm.incomingEvents <- slack.RTMEvent{
		Data: &slack.MessageEvent{
			Msg: slack.Msg{
				Text:    text,
				Channel: channel,
			},
		},
	}
	close(rtm.incomingEvents)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
