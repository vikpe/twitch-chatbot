package chatbot

import "github.com/gempir/go-twitch-irc/v3"

func IsBroadcaster(user twitch.User) bool {
	return hasBadge(user, "broadcaster")
}

func IsModerator(user twitch.User) bool {
	return hasBadge(user, "moderator") || IsBroadcaster(user)
}

func IsSubscriber(user twitch.User) bool {
	return hasBadge(user, "subscriber") || IsBroadcaster(user)
}

func hasBadge(user twitch.User, badge string) bool {
	if value, ok := user.Badges[badge]; ok {
		return value > 0
	}

	return false
}
