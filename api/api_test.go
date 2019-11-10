package api

import (
	"github.com/brianvoe/gofakeit"
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

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
