package chatbot_test

import (
	"testing"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/stretchr/testify/assert"
	chatbot "github.com/vikpe/twitch-chatbot"
)

func TestIsBroadcaster(t *testing.T) {
	t.Run("undefined value", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{}}
		assert.False(t, chatbot.IsBroadcaster(user))
	})

	t.Run("is not a broadcaster", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{"broadcaster": 0}}
		assert.False(t, chatbot.IsBroadcaster(user))
	})

	t.Run("is a broadcaster", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{"broadcaster": 1}}
		assert.True(t, chatbot.IsBroadcaster(user))
	})
}

func TestIsModerator(t *testing.T) {
	t.Run("undefined value", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{}}
		assert.False(t, chatbot.IsModerator(user))
	})

	t.Run("is not a moderator", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{"moderator": 0}}
		assert.False(t, chatbot.IsModerator(user))
	})

	t.Run("is a moderator", func(t *testing.T) {
		t.Run("normal moderator", func(t *testing.T) {
			user := twitch.User{Badges: map[string]int{"moderator": 1}}
			assert.True(t, chatbot.IsModerator(user))
		})

		t.Run("broadcaster", func(t *testing.T) {
			user := twitch.User{Badges: map[string]int{"broadcaster": 1}}
			assert.True(t, chatbot.IsModerator(user))
		})
	})
}

func TestIsSubscriber(t *testing.T) {
	t.Run("undefined value", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{}}
		assert.False(t, chatbot.IsSubscriber(user))
	})

	t.Run("is not a subscriber", func(t *testing.T) {
		user := twitch.User{Badges: map[string]int{"subscriber": 0}}
		assert.False(t, chatbot.IsSubscriber(user))
	})

	t.Run("is a subscriber", func(t *testing.T) {
		t.Run("normal subscriber", func(t *testing.T) {
			user := twitch.User{Badges: map[string]int{"subscriber": 1}}
			assert.True(t, chatbot.IsSubscriber(user))
		})

		t.Run("broadcaster", func(t *testing.T) {
			user := twitch.User{Badges: map[string]int{"broadcaster": 1}}
			assert.True(t, chatbot.IsSubscriber(user))
		})
	})
}
