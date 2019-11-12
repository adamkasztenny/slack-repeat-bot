package domain

import (
	"math/rand"
	"strings"
	"time"
)

func Repeat(word string) string {
	repetition := getNumberOfRepetitions()

	if len(word) < 3 {
		return strings.Repeat(word, repetition)
	}

	prefixToRepeat := string(word[0:2])
	repeatedPrefix := strings.Repeat(prefixToRepeat, repetition)
	return repeatedPrefix + word
}

func getNumberOfRepetitions() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8) + 2
}
