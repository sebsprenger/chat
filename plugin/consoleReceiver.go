package plugin

import (
	"fmt"
	"strings"

	"github.com/sebsprenger/chatterschool/shared"
)

type ConsoleReceiver struct {
	encryption bool
}

func (formatter ConsoleReceiver) FormatMessage(msg shared.Message) {
	output := msg.Text
	if formatter.encryption == true {
		output = formatter.decode(output)
	}
	fmt.Printf("%s: %s\n", msg.Sender, output)
}

func (formatter ConsoleReceiver) decode(input string) string {
	return strings.Map(decrypt, input)
}
