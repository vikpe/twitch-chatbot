package chatbot

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

type Chatbot struct {
	client           *twitch.Client
	channel          string
	commandHandlers  map[string]CommandHandler
	stopChan         chan os.Signal
	OnStarted        func()
	OnConnected      func()
	OnStopped        func(os.Signal)
	OnUnknownCommand CommandHandler
}

func NewChatbot(username string, oauth string, channel string, commandPrefix rune) *Chatbot {
	client := twitch.NewClient(username, oauth)
	client.Join(channel)

	bot := Chatbot{
		client:          client,
		channel:         channel,
		commandHandlers: make(map[string]CommandHandler, 0),
		OnStarted:       func() {},
		OnConnected:     func() {},
		OnStopped:       func(os.Signal) {},
	}

	bot.OnUnknownCommand = func(cmd Command, msg twitch.PrivateMessage) {
		replyMessage := fmt.Sprintf(`unknown command "%s". available commands: %s`, cmd.Name, bot.GetCommands(", "))
		bot.Reply(msg, replyMessage)
	}

	client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		if msg.Channel != channel {
			return
		}

		cmd, err := NewCommandFromMessage(commandPrefix, msg.Message)

		if err != nil {
			return
		}

		if cmdHandler, ok := bot.commandHandlers[cmd.Name]; ok {
			cmdHandler(cmd, msg)
		} else {
			bot.OnUnknownCommand(cmd, msg)
		}
	})

	return &bot
}

func (b *Chatbot) AddCommand(name string, handler CommandHandler) {
	b.commandHandlers[name] = handler
}

func (b *Chatbot) GetCommands(sep string) string {
	var commands []string

	for k := range b.commandHandlers {
		commands = append(commands, k)
	}

	sort.Strings(commands)

	return strings.Join(commands, sep)
}

func (b *Chatbot) Reply(msg twitch.PrivateMessage, replyText string) {
	b.client.Reply(msg.Channel, msg.ID, replyText)
}

func (b *Chatbot) Say(text string) {
	b.client.Say(b.channel, text)
}

func (b *Chatbot) Start() {
	b.OnStarted()

	b.stopChan = make(chan os.Signal, 1)
	signal.Notify(b.stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		b.client.OnConnect(func() {
			b.OnConnected()
		})
		b.client.Connect()
		defer b.client.Disconnect()
	}()
	sig := <-b.stopChan

	b.OnStopped(sig)
}

func (b *Chatbot) Stop() {
	if b.stopChan == nil {
		return
	}
	b.stopChan <- syscall.SIGINT
	time.Sleep(30 * time.Millisecond)
}
