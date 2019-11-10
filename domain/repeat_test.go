package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepeatTestSuite struct {
	suite.Suite
}

func (suite *RepeatTestSuite) TestRepeatingATwoLetterWord() {
	word := "go"
	result := Repeat(word)

	assert.Contains(suite.T(), result, word)
	assert.Contains(suite.T(), result, word+word)
}

func (suite *RepeatTestSuite) TestRepeatingAWordLongerThanTwoLetters() {
	word := "yeah"
	result := Repeat(word)

	assert.Contains(suite.T(), result, word)
	assert.Contains(suite.T(), result, "yeye")
}

func TestRepeatTestSuite(t *testing.T) {
	suite.Run(t, new(RepeatTestSuite))
}
