package plugin

import (
	"fmt"

	"github.com/sebsprenger/chatterschool/shared"
)

func NewConsoleDisplay(encryption Encryption) ConsoleDisplay {
	return ConsoleDisplay{
		encryption: encryption,
	}
}

type ConsoleDisplay struct {
	encryption Encryption
}

func (formatter ConsoleDisplay) FormatMessage(msg shared.Message) {
	output := formatter.encryption.Decrypt(msg.Text)
	fmt.Printf("%s: %s\n", msg.Sender, output)
}
