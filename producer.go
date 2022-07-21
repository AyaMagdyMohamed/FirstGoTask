package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

func main() {
	conn, err := stomp.Dial("tcp", "localhost:61616", nil)
	if err != nil {
		fmt.Println("---------err--------", err)
	}
	data := conn.Send("SampleQueue", "text/plain", []byte("Test"), nil)
	fmt.Println("---------data--------", data)
	testQueuedata := conn.Send("TestQueue", "", []byte("Test"), nil)
	fmt.Println("---------data--------", testQueuedata)
	conn.Disconnect()
}
