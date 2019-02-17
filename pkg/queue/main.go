package queue

import (
	"github.com/cskr/pubsub"
	event "github.com/valtyr/shake/pkg/event"
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

func ReadEvent(channel chan interface{}) event.Event {
	return (<-channel).(event.Event)
}
