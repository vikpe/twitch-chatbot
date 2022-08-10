package chatbot_test

import (
	"testing"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/stretchr/testify/assert"
	chatbot "github.com/vikpe/twitch-chatbot"
)

func TestChatbot_GetCommands(t *testing.T) {
	bot := chatbot.NewChatbot("foo", "oauth:bar", "baz", '!')
	bot.AddCommand("test1", func(cmd chatbot.Command, msg twitch.PrivateMessage) {})
	bot.AddCommand("test2", func(cmd chatbot.Command, msg twitch.PrivateMessage) {})

	assert.Equal(t, "test1, test2", bot.GetCommands(", "))
}
