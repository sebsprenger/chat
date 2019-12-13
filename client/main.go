package main

import (
	"flag"
	"log"

	"github.com/sebsprenger/chatclient/plugin"
	"github.com/sebsprenger/chatterschool/client"
)

var (
	ip   = flag.String("ip", "localhost", "server ip")
	port = flag.String("port", "9001", "server port")
	name = flag.String("name", "nobody", "name used for chat")
)

func main() {
	flag.Parse()

	consoleOutput := plugin.ConsoleReceiver{}

	consoleInput := ConsoleSender{
		scanner:          NewConsoleChatScanner(),
		messageFormatter: plugin.NewMessageFormatter(*name),
	}

	client := client.NewChatClient()
	err := client.Connect(*ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	client.ReceiveChatMessagenOn(consoleOutput)
	consoleInput.SendChatMessagesTo(&client)

}
