package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/go-stomp/stomp"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Aya", 0).Err()
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	// val, err := client.Get("name2").Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(val)
	conn, _ := stomp.Dial("tcp", "localhost:61616")
	sub, _ := conn.Subscribe("TestQueue2", stomp.AckAuto)
	msg := <-sub.C
	println(string(msg.Body))

	conn.Disconnect()
}
