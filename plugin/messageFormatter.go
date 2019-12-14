package plugin

import (
	"strings"

	"github.com/sebsprenger/chatterschool/shared"
)

type MessageFormatter struct {
	sender     string
	encryption bool
}

func NewMessageFormatter(senderName string, encryption bool) MessageFormatter {
	return MessageFormatter{
		sender:     senderName,
		encryption: encryption,
	}
}

func (formatter MessageFormatter) CreateMessage(input string) shared.Message {
	input = formatter.changeInput(input)
	if formatter.encryption == true {
		input = formatter.encode(input)
	}
	return formatter.buildMessage(input)
}

func (formatter MessageFormatter) buildMessage(input string) shared.Message {
	msg := shared.Message{
		Text:   input,
		Sender: formatter.sender,
	}
	return msg
}

func (formatter MessageFormatter) changeInput(input string) string {
	if input == "/shrug" {
		input = `¯\_(ツ)_/¯`
	}
	return input
}

func (formatter MessageFormatter) encode(input string) string {
	return strings.Map(encrypt, input)
}
