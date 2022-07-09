# twitch-chatbot

> Twitch chatbot in Go (Golang)

Simple chatbot interface based on [github.com/gempir/go-twitch-irc](https://github.com/gempir/go-twitch-irc)

## Install

```shell
go get github.com/vikpe/twitch-chatbot
```

## Generate chatbot oauth token

* [Twitch Chat OAuth Password Generator](https://twitchapps.com/tmi/)

## Usage

```go

package bot_test

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/vikpe/twitch-chatbot"
)

func RunChatbot() {
	// init
	username := "bot_username"
	oauth := "oauth:bbbbbbbbbbbbb"
	channel := "channel_name"
	commandPrefix := '!'

	myBot := chatbot.NewChatbot(username, oauth, channel, commandPrefix)

	// event callbacks
	myBot.OnStarted = func() { fmt.Println("chatbot started") }
	myBot.OnConnected = func() { fmt.Println("chatbot connected") }
	myBot.OnStopped = func(sig os.Signal) {
		fmt.Println(fmt.Sprintf("chatbot stopped (%s)", sig))
	}

	// command handlers
	myBot.AddCommand("hello", func(cmd chatbot.Command, msg twitch.PrivateMessage) {
		myBot.Reply(msg, "world!")
	})

	myBot.AddCommand("test", func(cmd chatbot.Command, msg twitch.PrivateMessage) {
		myBot.Say(fmt.Sprintf("%s called the test command using args %s", msg.User.Name, cmd.ArgsToString()))
	})

	myBot.AddCommand("mod_only", func(cmd chatbot.Command, msg twitch.PrivateMessage) {
		if !chatbot.IsBroadcaster(msg.User) {
			myBot.Reply(msg, "mod_only is only allowed by moderators.")
			return
		}

		myBot.Say(fmt.Sprintf("%s called the mod_only command", msg.User.Name))
	})

	myBot.Start() // blocking operation
}
```
