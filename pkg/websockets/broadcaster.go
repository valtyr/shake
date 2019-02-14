package websockets

import (
	"encoding/json"

	queue "github.com/valtyr/shake/pkg/queue"
)

func WebsocketBroadcaster(clientManager *ClientManager) {
	channel := queue.SubscribeToEvents()
	for {
		event := (<-channel)
		encoded, _ := json.Marshal(event)
		clientManager.broadcast <- encoded
	}
}
