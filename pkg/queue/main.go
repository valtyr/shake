package queue

import (
	"github.com/cskr/pubsub"
)

var queue *pubsub.PubSub

func CreateQueue() {
	queue = pubsub.New(1000)
}

func PublishEvent(event interface{}) {
	queue.Pub(event, "events")
}

func SubscribeToEvents() chan interface{} {
	return queue.Sub("events")
}
