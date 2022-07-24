package main

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/go-stomp/stomp"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Id   string `json:"id"`
}

func main() {
	// get JSON from amq
	conn, _ := stomp.Dial("tcp", "localhost:61616")
	sub, err := conn.Subscribe("SampleQueue", stomp.AckAuto)
	if err != nil {
		fmt.Println("Coudnl't receive from queue")
	}
	msg := <-sub.C
	println(string(msg.Body))

	student := Student{}
	json.Unmarshal(msg.Body, &student)

	fmt.Println("studentEncoded ", student)
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// set JSON object to redis
	err = client.Set(student.Id, msg.Body, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	conn.Disconnect()
}
