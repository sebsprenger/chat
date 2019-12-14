package plugin

import (
	"github.com/sebsprenger/chatterschool/shared"
)

type MessagePreparer struct {
	sender     string
	encryption Encryption
}

func NewMessagePreparer(senderName string, encryption Encryption) MessagePreparer {
	return MessagePreparer{
		sender:     senderName,
		encryption: encryption,
	}
}

func (formatter MessagePreparer) CreateMessage(input string) shared.Message {
	input = formatter.changeInput(input)
	input = formatter.encryption.Encrypt(input)

	return formatter.buildMessage(input)
}

func (formatter MessagePreparer) buildMessage(input string) shared.Message {
	msg := shared.Message{
		Text:   input,
		Sender: formatter.sender,
	}
	return msg
}

func (formatter MessagePreparer) changeInput(input string) string {
	if input == "/shrug" {
		input = `¯\_(ツ)_/¯`
	}
	if input == "/shrek" {
		input = `shrek ist einer meiner Lieblingsfilme`
	}
	return input
}
