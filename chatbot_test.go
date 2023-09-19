package chatbot_test

import (
	"testing"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/stretchr/testify/assert"
	chatbot "github.com/vikpe/twitch-chatbot"
)

func TestChatbot_GetCommands(t *testing.T) {
	bot := chatbot.NewChatbot("foo", "oauth:bar", "baz", '!')
	bot.AddCommand("gamma", func(cmd chatbot.Command, msg twitch.PrivateMessage) {})
	bot.AddCommand("alpha", func(cmd chatbot.Command, msg twitch.PrivateMessage) {})
	bot.AddCommand("beta", func(cmd chatbot.Command, msg twitch.PrivateMessage) {})

	assert.Equal(t, "alpha, beta, gamma", bot.GetCommands(", "))
}
