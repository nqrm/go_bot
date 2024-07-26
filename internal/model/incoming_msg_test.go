package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	model := New(nil)

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}
