package main

import (
	"flag"
	"log"

	"github.com/sebsprenger/chatclient/plugin"
	"github.com/sebsprenger/chatterschool/client"
)

var (
	ip         = flag.String("ip", "localhost", "server ip")
	port       = flag.String("port", "9001", "server port")
	name       = flag.String("name", "nobody", "name used for chat")
	encryption = flag.Bool("enc", false, "use encryption")
	caesar     = flag.Int("caesar", 0, "defines the key for caesar encryption")
)

func main() {
	flag.Parse()

	encryptionAlgorithm := plugin.NewCaesar(*encryption, *caesar)

	consoleInput := ConsoleSender{
		scanner:         NewConsoleChatScanner(),
		messagePreparer: plugin.NewMessagePreparer(*name, encryptionAlgorithm),
	}
	consoleOutput := plugin.NewConsoleDisplay(encryptionAlgorithm)

	client := client.NewChatClient()
	err := client.Connect(*ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	client.ReceiveChatMessagenOn(consoleOutput)
	consoleInput.SendChatMessagesTo(&client)
}
