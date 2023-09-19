package chatbot

import (
	"errors"
	"strings"
	"unicode"

	"github.com/gempir/go-twitch-irc/v4"
)

type Command struct {
	Name string
	Args []string
}

type CommandHandler func(cmd Command, msg twitch.PrivateMessage)

func NewCommand(name string, args ...string) Command {
	return Command{
		Name: name,
		Args: args,
	}
}

func NewCommandFromMessage(prefix rune, text string) (Command, error) {
	if !IsCommand(prefix, text) {
		return Command{}, errors.New("unable to parse command call")
	}
	txt := strings.TrimLeft(text, " ")
	txt = strings.ToLower(txt[1:])
	parts := strings.FieldsFunc(txt, unicode.IsSpace)
	name := parts[0]
	args := parts[1:]
	return NewCommand(name, args...), nil
}

func (c Command) ArgsToString() string {
	return strings.Join(c.Args, " ")
}

func IsCommand(prefix rune, text string) bool {
	txt := strings.TrimLeft(text, " ")

	if len(txt) < 2 || 0 != strings.IndexRune(txt, prefix) {
		return false
	}

	parts := strings.FieldsFunc(txt[1:], unicode.IsSpace)

	if 0 == len(parts) {
		return false
	}

	firstRune := rune(parts[0][0])
	return unicode.IsLetter(firstRune)
}
