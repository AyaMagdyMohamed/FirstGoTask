package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-stomp/stomp"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   string `json:"id"`
}

func main() {

	conn, err := stomp.Dial("tcp", "localhost:61616", nil)
	if err != nil {
		fmt.Println("---------err--------", err)
	}
	student := Student{Name: "Ahmed", Age: 25, Id: "1"}
	studentEncoded, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	data := conn.Send("SampleQueue", "text/plain", studentEncoded, nil)
	fmt.Println("---------data--------", data)
	conn.Disconnect()
}
