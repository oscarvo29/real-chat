package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oscarvo29/real-chat/chat-listener/utils"
)

type Config struct {
}

func main() {
	ctx := context.Background()

	pool := GetConnection(utils.GetEnvValue("DSN"))
	fmt.Println("Connected to Postgres :D ")
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = conn.Exec(ctx, "listen new_message")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Release()

	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("notification payload: ", notification.Payload)
	}
}
