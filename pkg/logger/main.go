package logger

import (
	"fmt"

	queue "github.com/valtyr/shake/pkg/queue"
)

func Logger() {
	channel := queue.SubscribeToEvents()
	for {
		event := queue.ReadEvent(channel)
		fmt.Println("Event received: " + event.Type)
	}
}
