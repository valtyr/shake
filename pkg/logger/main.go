package logger

import (
	"fmt"

	queue "github.com/valtyr/shake/pkg/queue"
)

func Logger() {
	channel := queue.SubscribeToEvents()
	for {
		(<-channel)
		fmt.Println("Event received")
	}
}
