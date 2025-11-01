package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Sender  string
	Content string
}

type Reply struct {
	History []string
}

func main() {
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Start chatting! Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		msg := Message{Sender: name, Content: text}
		var reply Reply

		// Call the server's SendMessage method
		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			log.Println("Error calling remote procedure:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range reply.History {
			fmt.Println(m)
		}
		fmt.Println("--------------------")
	}
}